-- This script inserts 2 sample frameworks into the "frameworkstable"
INSERT INTO public.frameworks (
        "frameworks_uuid",
        "name",
        "created_at",
        "updated_at"
    )
VALUES (
        gen_random_uuid(),
        'MPA',
        current_timestamp,
        current_timestamp
    ),
    (
        gen_random_uuid(),
        'CIS',
        current_timestamp,
        current_timestamp
    );