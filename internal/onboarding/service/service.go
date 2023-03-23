package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	companyEntities "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/company/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/onboarding/repository"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/cache"
)

// Service stores business logic
type Service interface {
	UpdateOnboardingStatus(ctx context.Context, onboardingStep int, companyUuid, updated_by uuid.UUID)
}

type service struct {
	repo  repository.Repository
	cache cache.Client
}

func New(repo repository.Repository, cache cache.Client) Service {
	return &service{
		repo:  repo,
		cache: cache,
	}
}

func (s *service) UpdateOnboardingStatus(ctx context.Context, onboardingStep int, companyUuid, updated_by uuid.UUID) {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go s.updateOnboardingStatus(ctx, onboardingStep, companyUuid, updated_by, wg)
	wg.Wait()
}

func (s *service) cacheGetOnboardingStatus(ctx context.Context, companyUuid uuid.UUID) (*companyEntities.OnboardingGroup, error) {
	var cached_status *companyEntities.OnboardingGroup
	var err error

	data, _ := s.cache.Get(ctx, fmt.Sprintf("%s_%s", entities.OnboardingStatusCacheKey, companyUuid.String()))

	if len(data) != 0 {
		err := json.Unmarshal([]byte(data), &cached_status)
		if err != nil {
			return nil, err
		}
	}
	// we didn't find anything in cache, lets update it
	if cached_status == nil {
		cached_status, err = s.cacheSetOnboardingStatus(ctx, companyUuid)
		if err != nil {
			log.Println(err)
		}
	}

	return cached_status, nil
}

func (s *service) cacheSetOnboardingStatus(ctx context.Context, companyUuid uuid.UUID) (*companyEntities.OnboardingGroup, error) {
	log.Println("Caching Onboarding Status")

	// call db to get latest statuses
	status, err := s.repo.GetOnboardingStatus(ctx, companyUuid)
	if err != nil {
		return nil, err
	}

	out, err := json.Marshal(status)
	if err != nil {
		return nil, err
	}

	// update cache with latest onboarding data
	err = s.cache.Set(ctx, fmt.Sprintf("%s_%s", entities.OnboardingStatusCacheKey, companyUuid.String()), string(out))
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (s *service) updateOnboardingStatus(ctx context.Context, onboardingStep int, companyUuid, updated_by uuid.UUID, wg *sync.WaitGroup) {
	defer wg.Done()

	var updatedOnboardingSteps companyEntities.OnboardingGroup
	var updatedStep companyEntities.Onboarding

	prevStatus, err := s.cacheGetOnboardingStatus(ctx, companyUuid)
	if err != nil {
		log.Println(err)
	}

	if prevStatus != nil {
		for _, step := range *prevStatus {
			if step.Name == entities.OnboardinStepToName[onboardingStep] && step.Status == entities.OnboardingStatusComplete {
				// onboarding step is already completed, skip everything except security awareness training step
				if !(entities.SecurityAwarenessTrainingStep == entities.OnboardinStepToName[onboardingStep]) {
					log.Println("Step is already complete")
					return
				}

			}
		}

		for _, step := range *prevStatus {
			updatedStep = step
			if step.Name == entities.OnboardinStepToName[onboardingStep] {
				updatedStep.Status = entities.OnboardingStatusComplete
				updatedStep.UpdatedBy = updated_by.String()
				updatedStep.UpdatedAt = time.Now().Format(time.RFC3339)
			}
			updatedOnboardingSteps = append(updatedOnboardingSteps, updatedStep)
		}
	}

	log.Println("Updating onboarding step status")
	err = s.repo.UpdateOnboardingStatus(ctx, &updatedOnboardingSteps, companyUuid, updated_by)
	if err != nil {
		log.Println(err)
	}

	// once a step is completed, update its cached value as well
	_, err = s.cacheSetOnboardingStatus(ctx, companyUuid)
	if err != nil {
		log.Println(err)
	}
}
