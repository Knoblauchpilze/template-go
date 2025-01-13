-- https://dba.stackexchange.com/questions/117109/how-to-manage-default-privileges-for-users-on-a-database-vs-schema/117661#117661
CREATE DATABASE db_products OWNER products_admin;
REVOKE ALL ON DATABASE db_products FROM public;

GRANT CONNECT ON DATABASE db_products TO products_user;

\connect db_products

CREATE SCHEMA products_schema AUTHORIZATION products_admin;

SET search_path = products_schema;

ALTER ROLE products_admin IN DATABASE db_products SET search_path = products_schema;
ALTER ROLE products_manager IN DATABASE db_products SET search_path = products_schema;
ALTER ROLE products_user IN DATABASE db_products SET search_path = products_schema;

GRANT USAGE  ON SCHEMA products_schema TO products_user;
GRANT CREATE ON SCHEMA products_schema TO products_admin;

ALTER DEFAULT PRIVILEGES FOR ROLE products_admin
GRANT SELECT ON TABLES TO products_user;

ALTER DEFAULT PRIVILEGES FOR ROLE products_admin
GRANT INSERT, UPDATE, DELETE ON TABLES TO products_manager;
