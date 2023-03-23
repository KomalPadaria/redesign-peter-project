-- +migrate Up
ALTER TABLE public.policy_histories DROP CONSTRAINT IF EXISTS fk_policies;
ALTER TABLE public.policy_histories ADD CONSTRAINT fk_policies FOREIGN KEY (policy_uuid) REFERENCES public.policies(policy_uuid) ON DELETE CASCADE;

-- +migrate Down
ALTER TABLE public.policy_histories DROP CONSTRAINT IF EXISTS fk_policies;
ALTER TABLE public.policy_histories ADD CONSTRAINT fk_policies FOREIGN KEY (policy_uuid) REFERENCES public.policies(policy_uuid);