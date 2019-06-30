REVOKE ALL PRIVILEGES ON PROCEDURE Dare.get_all_tasks FROM dare_api;
REVOKE ALL PRIVILEGES ON PROCEDURE Dare.get_task FROM dare_api;
REVOKE ALL PRIVILEGES ON PROCEDURE Dare.update_task FROM dare_api;
REVOKE ALL PRIVILEGES ON PROCEDURE Dare.create_task FROM dare_api;
REVOKE ALL PRIVILEGES ON PROCEDURE Dare.delete_task FROM dare_api;

DROP PROCEDURE Dare.get_all_tasks;
DROP PROCEDURE Dare.get_task;
DROP PROCEDURE Dare.create_task;
DROP PROCEDURE Dare.update_task;
DROP PROCEDURE Dare.delete_task;

