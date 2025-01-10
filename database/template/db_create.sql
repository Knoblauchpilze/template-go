-- https://dba.stackexchange.com/questions/117109/how-to-manage-default-privileges-for-users-on-a-database-vs-schema/117661#117661
CREATE DATABASE db_template_service OWNER template_service_admin;
REVOKE ALL ON DATABASE db_template_service FROM public;

GRANT CONNECT ON DATABASE db_template_service TO template_service_user;

\connect db_template_service

CREATE SCHEMA template_service_schema AUTHORIZATION user_service_admin;

SET search_path = template_service_schema;

ALTER ROLE template_service_admin IN DATABASE db_template_service SET search_path = template_service_schema;
ALTER ROLE template_service_manager IN DATABASE db_template_service SET search_path = template_service_schema;
ALTER ROLE template_service_user IN DATABASE db_template_service SET search_path = template_service_schema;

GRANT USAGE  ON SCHEMA template_service_schema TO template_service_user;
GRANT CREATE ON SCHEMA template_service_schema TO template_service_admin;

ALTER DEFAULT PRIVILEGES FOR ROLE template_service_admin
GRANT SELECT ON TABLES TO template_service_user;

ALTER DEFAULT PRIVILEGES FOR ROLE template_service_admin
GRANT INSERT, UPDATE, DELETE ON TABLES TO template_service_manager;
