-- Create "organizations" table
CREATE TABLE `organizations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `organization_name` longtext NOT NULL,
  `organization_code` varchar(191) NOT NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uni_organizations_organization_code` (`organization_code`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "customers" table
CREATE TABLE `customers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `organization_id` bigint unsigned NOT NULL,
  `name` longtext NOT NULL,
  `unit_price` bigint NULL,
  `start_date` date NULL,
  `end_date` date NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_customers_deleted_at` (`deleted_at`),
  INDEX `idx_customers_organization_id` (`organization_id`),
  CONSTRAINT `fk_organizations_customers` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "projects" table
CREATE TABLE `projects` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `organization_id` bigint unsigned NOT NULL,
  `customer_id` bigint unsigned NOT NULL,
  `name` longtext NOT NULL,
  `unit_price` bigint NULL,
  `start_date` date NULL,
  `end_date` date NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_projects_customer_id` (`customer_id`),
  INDEX `idx_projects_deleted_at` (`deleted_at`),
  INDEX `idx_projects_organization_id` (`organization_id`),
  CONSTRAINT `fk_customers_projects` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT `fk_organizations_projects` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "budgets" table
CREATE TABLE `budgets` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `organization_id` bigint unsigned NOT NULL,
  `project_id` bigint unsigned NOT NULL,
  `amount` bigint NOT NULL,
  `memo` text NULL,
  `start_date` date NULL,
  `end_date` date NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_budgets_organization_id` (`organization_id`),
  INDEX `idx_budgets_project_id` (`project_id`),
  CONSTRAINT `fk_organizations_budgets` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT `fk_projects_budgets` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "users" table
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `organization_id` bigint unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `role` enum('admin','member') NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_users_deleted_at` (`deleted_at`),
  INDEX `idx_users_organization_id` (`organization_id`),
  UNIQUE INDEX `uni_users_email` (`email`),
  CONSTRAINT `fk_organizations_users` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "timers" table
CREATE TABLE `timers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `organization_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `project_id` bigint unsigned NULL,
  `title` longtext NULL,
  `memo` longtext NULL,
  `start_at` datetime(3) NOT NULL,
  `end_at` datetime(3) NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_timers_end_at` (`end_at`),
  INDEX `idx_timers_organization_id` (`organization_id`),
  INDEX `idx_timers_project_id` (`project_id`),
  INDEX `idx_timers_start_at` (`start_at`),
  INDEX `idx_timers_user_id` (`user_id`),
  CONSTRAINT `fk_organizations_timer` FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT `fk_projects_timer` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT `fk_users_timer` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
