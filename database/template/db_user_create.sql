
-- https://stackoverflow.com/questions/72985242/securely-create-role-in-postgres-using-sql-script-and-environment-variables
CREATE USER template_service_admin WITH CREATEDB PASSWORD :'admin_password';
CREATE USER template_service_manager WITH PASSWORD :'manager_password';
CREATE USER template_service_user WITH PASSWORD :'user_password';

GRANT template_service_user TO template_service_manager;
GRANT template_service_manager TO template_service_admin;
