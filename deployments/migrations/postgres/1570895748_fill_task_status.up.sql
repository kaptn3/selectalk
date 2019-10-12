INSERT INTO task_status
  (id, updated_at, created_at, status)
VALUES (DEFAULT, now(), now(), 'InProgress'), (DEFAULT, now(), now(), 'Done');