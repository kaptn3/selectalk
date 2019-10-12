INSERT INTO task
  (id, created_at, updated_at, "name", score, expiration_date, user_id, manager_id, task_status_id, task_priority_id)
VALUES
  (DEFAULT, now(), now(), 'Подключить сервера', 100, now() + INTERVAL '1 WEEK', 1, 2, 1, 1),
  (DEFAULT, now(), now(), 'Проверить исправность облачного хранилища', 50, now() + INTERVAL '2 WEEK', 1, 2, 1, 2),
  (DEFAULT, now(), now(), 'Оформить техническое задание', 150, now() + INTERVAL '3 WEEK', 1, 2, 1, 3);