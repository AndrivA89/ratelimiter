Rate Limiter в виде пакета
==
1) на входе принимает канал с задачами (просто числа типа int) и запускает задачи параллельно
2) имеет два ограничения:
- максимальное количество одновременных задач
- максимальное количество задач в течении минуты
3) оба ограничения настраиваются в момент инициализации (через переменные окружения)

`RATE_LIMITER_MAX_TASKS_IN_MINUTE` - максимальное количество задач в минуту

`RATE_LIMITER_MAX_WORKERS` - максимальное количество одновременных задач