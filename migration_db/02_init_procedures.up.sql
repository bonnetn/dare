CREATE PROCEDURE Dare.get_all_tasks()
  BEGIN
    SELECT uuid, name, content, version FROM Dare.Tasks;
  END;

CREATE PROCEDURE Dare.get_task(IN uuid_p CHAR(36))
  BEGIN
    SELECT uuid, name, content, version FROM Dare.Tasks WHERE uuid=uuid_p;
  END;

CREATE PROCEDURE Dare.create_task(IN uuid_p CHAR(36), IN version_p BIGINT, IN name_p TEXT, IN content_p JSON)
  BEGIN
    INSERT IGNORE INTO Dare.Tasks (uuid, name, content, version)
      VALUES (uuid_p, name_p, content_p, version_p);

    CALL get_task(uuid_p);
  END;

CREATE PROCEDURE Dare.update_task(IN uuid_p CHAR(36), IN version_p BIGINT, IN name_p TEXT, IN content_p JSON)
  BEGIN
    UPDATE IGNORE Dare.Tasks
      SET uuid=uuid_p, name=name_p, content=content_p, version=version_p
      WHERE uuid=uuid_p AND version=version_p-1;

    CALL get_task(uuid_p);
  END;

CREATE PROCEDURE Dare.delete_task(IN uuid_p CHAR(36))
  BEGIN
    DELETE FROM Dare.Tasks WHERE uuid=uuid_p;
  END;

GRANT EXECUTE ON PROCEDURE Dare.get_all_tasks TO dare_api;
GRANT EXECUTE ON PROCEDURE Dare.get_task TO dare_api;
GRANT EXECUTE ON PROCEDURE Dare.update_task TO dare_api;
GRANT EXECUTE ON PROCEDURE Dare.create_task TO dare_api;
GRANT EXECUTE ON PROCEDURE Dare.delete_task TO dare_api;
