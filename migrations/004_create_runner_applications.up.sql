CREATE TABLE runner_applications (
  id                    UUID        PRIMARY KEY DEFAULT uuid_generate_v4(),
  name                  TEXT        NOT NULL,
  phone                 TEXT        NOT NULL,
  ic_number             TEXT        UNIQUE NOT NULL,
  vehicle_type          TEXT        NOT NULL CHECK (vehicle_type IN ('motorbike','car','bicycle')),
  plate_number          TEXT        NOT NULL,
  pet_experience        TEXT[],
  comfortable_with_live_pets BOOLEAN,
  consent_acknowledged  BOOLEAN     NOT NULL,
  status                TEXT        NOT NULL DEFAULT 'pending_review',
  submitted_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  reviewed_at           TIMESTAMPTZ,
  reviewer_user_id      UUID
);
