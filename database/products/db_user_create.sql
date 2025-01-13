
-- https://stackoverflow.com/questions/72985242/securely-create-role-in-postgres-using-sql-script-and-environment-variables
CREATE USER products_admin WITH CREATEDB PASSWORD :'admin_password';
CREATE USER products_manager WITH PASSWORD :'manager_password';
CREATE USER products_user WITH PASSWORD :'user_password';

GRANT products_user TO products_manager;
GRANT products_manager TO products_admin;
