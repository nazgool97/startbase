-- Добавляем роль и поля для сброса пароля
ALTER TABLE users
ADD COLUMN IF NOT EXISTS role VARCHAR(16) NOT NULL DEFAULT 'user',
ADD COLUMN IF NOT EXISTS reset_token VARCHAR(64),
ADD COLUMN IF NOT EXISTS reset_expires TIMESTAMPTZ;