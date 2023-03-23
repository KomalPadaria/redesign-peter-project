package service

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/knowbe4"
	knowbe4Entity "github.com/nurdsoft/redesign-grp-trust-portal-api/internal/knowbe4/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/internal/securityawareness/entities"
	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/date"
)

type Service interface {
	GetPhishingDetails(ctx context.Context, companyUuid, userUuid *uuid.UUID) (*entities.GetPhishingDetailsResponse, error)
	GetTrainingDetails(ctx context.Context, companyUuid, userUuid *uuid.UUID) (*entities.GetTrainingDetailsResponse, error)
}

func New(knowBe4Client knowbe4.Client) Service {
	return &service{knowBe4Client}
}

type service struct {
	knowBe4Client knowbe4.Client
}

func (s *service) GetTrainingDetails(ctx context.Context, companyUuid, userUuid *uuid.UUID) (*entities.GetTrainingDetailsResponse, error) {
	now := time.Now()

	currentQuarterName, currentStartDate, currentEndDate, err := s.currentQuarterDetails(now)
	if err != nil {
		return nil, err
	}

	nextQuarterName, nextStartDate, nextEndDate, err := s.nextQuarterDetails(now)
	if err != nil {
		return nil, err
	}

	enrollments, err := s.knowBe4Client.GetAllTrainingEnrollments(ctx, *companyUuid)
	if err != nil {
		return nil, err
	}

	var notStarted, inProgress, completed, passed, pastDue, totalEnrollments int
	for _, e := range enrollments {
		if date.InTimeSpan(currentStartDate, currentEndDate, e.EnrollmentDate) {
			totalEnrollments++
			switch e.Status {
			case "Not Started":
				notStarted++
			case "In Progress":
				inProgress++
			case "Completed":
				completed++
			case "Passed":
				passed++
			case "Past Due":
				pastDue++
			}
		}
	}

	res := &entities.GetTrainingDetailsResponse{
		CurrentNextCampaign: entities.CurrentNextCampaign{
			Current: entities.Campaign{
				Name:      currentQuarterName,
				StartDate: currentStartDate,
				EndDate:   currentEndDate,
			},
			Next: entities.Campaign{
				Name:      nextQuarterName,
				StartDate: nextStartDate,
				EndDate:   nextEndDate,
			},
		},
		TrainingStat: entities.TrainingStat{
			NotStarted:       notStarted,
			InProgress:       inProgress,
			Completed:        completed,
			Passed:           passed,
			PastDue:          pastDue,
			TotalEnrollments: totalEnrollments,
		},
	}

	return res, nil
}

func (s *service) GetPhishingDetails(ctx context.Context, companyUuid, userUuid *uuid.UUID) (*entities.GetPhishingDetailsResponse, error) {
	now := time.Now()

	currentQuarterName, currentStartDate, currentEndDate, err := s.currentQuarterDetails(now)
	if err != nil {
		return nil, err
	}

	nextQuarterName, nextStartDate, nextEndDate, err := s.nextQuarterDetails(now)
	if err != nil {
		return nil, err
	}

	psts, err := s.knowBe4Client.GetAllPhishingSecurityTests(ctx, *companyUuid)
	if err != nil {
		return nil, err
	}
	recipientResults := make([]knowbe4Entity.RecipientResult, 0)
	for _, p := range psts {
		// check whether campaign/test start date falls under current quarter start and end date
		if date.InTimeSpan(currentStartDate, currentEndDate, p.StartedAt) {
			rrs, err := s.knowBe4Client.GetAllRecipientResults(ctx, *companyUuid, p.PstID)
			if err != nil {
				return nil, err
			}
			recipientResults = append(recipientResults, rrs...)
		}
	}

	userMap := make(map[int]knowbe4Entity.User)
	userClickCountMap := make(map[int]int)
	var passCount, failCount int

	for _, recipientResult := range recipientResults {
		userMap[recipientResult.User.ID] = recipientResult.User
		if !recipientResult.ClickedAt.IsZero() {
			userClickCountMap[recipientResult.User.ID]++
			failCount++
		} else {
			passCount++
		}
	}

	keys := make([]int, 0, len(userClickCountMap))

	for key := range userClickCountMap {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return userClickCountMap[keys[i]] > userClickCountMap[keys[j]]
	})
	clickers := make([]entities.Clicker, 0)

	topClickerFilterCount := 10

	for i, k := range keys {
		if i <= topClickerFilterCount-1 {
			user := userMap[k]
			c := entities.Clicker{
				FullName:     fmt.Sprintf("%s %s", user.FirstName, user.LastName),
				Email:        user.Email,
				ClickedCount: userClickCountMap[k],
			}
			clickers = append(clickers, c)
		}
	}

	phishing := &entities.GetPhishingDetailsResponse{
		CurrentNextCampaign: entities.CurrentNextCampaign{
			Current: entities.Campaign{
				Name:      currentQuarterName,
				StartDate: currentStartDate,
				EndDate:   currentEndDate,
			},
			Next: entities.Campaign{
				Name:      nextQuarterName,
				StartDate: nextStartDate,
				EndDate:   nextEndDate,
			},
		},
		PhishingStat: entities.PhishingStat{
			Passed:           passCount,
			Failed:           failCount,
			TotalParticipant: len(recipientResults),
		},
		TopClickers: clickers,
	}

	return phishing, nil
}

func (s *service) currentQuarterDetails(t time.Time) (string, time.Time, time.Time, error) {
	sd, ed, err := date.QuarterStartDateAndEndDateByTime(t)
	if err != nil {
		return "", time.Time{}, time.Time{}, err
	}

	quarter := date.QuarterOf(int(t.Month()))
	year := t.Year()
	quarterName := fmt.Sprintf("Q%d %d", quarter, year)

	return quarterName, sd, ed, nil
}

func (s *service) nextQuarterDetails(t time.Time) (string, time.Time, time.Time, error) {
	/* next quarter is after 3 months
	Ref:
		January, February, and March (Q1)
		April, May, and June (Q2)
		July, August, and September (Q3)
		October, November, and December (Q4)
	*/
	t = t.AddDate(0, 3, 0)
	sd, ed, err := date.NextQuarterStartDateAndEndDateByTime(t)
	if err != nil {
		return "", time.Time{}, time.Time{}, err
	}

	quarter := date.QuarterOf(int(t.Month()))
	year := t.Year()
	if quarter == 4 {
		year += 1
		quarter = 1
	}
	quarterName := fmt.Sprintf("Q%d %d", quarter, year)

	return quarterName, sd, ed, nil
}
