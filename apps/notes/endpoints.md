#### Heating Service
* __GET /api/v1/heating-systems__  
Описание: Получить список всех систем отопления (с фильтрами по house_id или user_id).  
Параметры: Query (например, ?house_id=123).  
Ответ: Массив объектов с id, name, status и т.д.

* __GET /api/v1/heating-systems/{id}__  
Описание: Получить детали конкретной системы отопления.  
Параметры: Path (id системы).  
Ответ: Объект с полными атрибутами (id, name, house_id, status, created_at и т.д.).

* __POST /api/v1/heating-systems__  
Описание: Добавить новую систему отопления.  
Тело запроса: JSON с name, house_id, status.  
Ответ: Созданный объект с id.  

* __PUT /api/v1/heating-systems/{id}__  
Описание: Обновить существующую систему отопления (например, изменить статус или имя).  
Тело запроса: JSON с обновляемыми полями.  
Ответ: Обновленный объект.  

* __DELETE /api/v1/heating-systems/{id}__  
Описание: Удалить систему отопления.  
Ответ: 204 No Content.  

#### Heating Monitoring Service
* __GET /api/v1/heating-monitoring/devices__  
Описание: Получить список отопительных устройств с текущим статусом (температура, состояние).  
Параметры: Query (например, ?house_id=123).  
Ответ: Массив объектов с id, current_temperature, status, last_reading.

* __GET /api/v1/heating-monitoring/devices/{id}/telemetry__  
Описание: Получить текущую телеметрию для конкретного устройства.  
Параметры: Path (id устройства).  
Ответ: Объект с timestamp, data (JSON с температурой и др.).  

* __GET /api/v1/heating-monitoring/devices/{id}/history__  
Описание: Получить историю телеметрии (например, за последние 24 часа).  
Параметры: Query (например, ?start=2023-01-01T00:00:00Z&end=2023-01-02T00:00:00Z).  
Ответ: Массив записей с timestamp, data.  

* __POST /api/v1/heating-monitoring/devices/{id}/alert__  
Описание: Отправить оповещение о проблеме (например, низкая температура).  
Тело запроса: JSON с reason, message.  
Ответ: 201 Created.  

#### Heating Schedule Service
* __GET /api/v1/heating-schedules__  
Описание: Получить список расписаний для пользователя или дома.  
Параметры: Query (например, ?heating_system_id=456).  
Ответ: Массив объектов с id, schedule (JSON), active.  

* __GET /api/v1/heating-schedules/{id}__  
Описание: Получить детали конкретного расписания. 
Параметры: Path (id расписания).  
Ответ: Объект с полными атрибутами.  

* __POST /api/v1/heating-schedules__  
Описание: Создать новое расписание (например, на основе времени и температуры).  
Тело запроса: JSON с heating_system_id, schedule (JSON с периодами и целями температуры), active.  
Ответ: Созданный объект с id.  

* __PUT /api/v1/heating-schedules/{id}__  
Описание: Обновить расписание (изменить активность или настройки).  
Тело запроса: JSON с обновляемыми полями.  
Ответ: Обновленный объект.  

* __DELETE /api/v1/heating-schedules/{id}__  
Описание: Удалить расписание.  
Ответ: 204 No Content.  

#### Lighting Service
* __GET /api/v1/lighting-devices__  
Описание: Получить список устройств освещения.  
Параметры: Query (например, ?house_id=123).  
Ответ: Массив объектов с id, status, brightness.  

* __GET /api/v1/lighting-devices/{id}__  
Описание: Получить детали конкретного устройства.  
Параметры: Path (id устройства).  
Ответ: Объект с полными атрибутами.  

* __POST /api/v1/lighting-devices__  
Описание: Добавить новое устройство освещения.  
Тело запроса: JSON с lighting_system_id, serial_number, status.  
Ответ: Созданный объект с id.  

* __PUT /api/v1/lighting-devices/{id}__  
Описание: Обновить устройство (например, изменить статус или яркость).  
Тело запроса: JSON с обновляемыми полями.  
Ответ: Обновленный объект.  

* __DELETE /api/v1/lighting-devices/{id}__  
Описание: Удалить устройство.  
Ответ: 204 No Content.  

#### Lighting Monitoring Service
* __GET /api/v1/lighting-monitoring/devices__  
Описание: Получить список устройств с текущим состоянием (яркость, статус).  
Параметры: Query (например, ?house_id=123).  
Ответ: Массив объектов с id, current_brightness, status.  

* __GET /api/v1/lighting-monitoring/devices/{id}/telemetry__  
Описание: Получить текущую телеметрию для устройства.  
Параметры: Path (id устройства).  
Ответ: Объект с timestamp, data (JSON с яркостью и др.).  

* __GET /api/v1/lighting-monitoring/devices/{id}/history__  
Описание: Получить историю телеметрии.  
Параметры: Query (например, ?start=...&end=...).  
Ответ: Массив записей.  

* __POST /api/v1/lighting-monitoring/devices/{id}/alert__  
Описание: Отправить оповещение о проблеме (например, сбой света).  
Тело запроса: JSON с reason, message.  
Ответ: 201 Created.  

#### Lighting Schedule Service
* __GET /api/v1/lighting-scenarios__  
Описание: Получить список сценариев.  
Параметры: Query (например, ?lighting_system_id=456).  
Ответ: Массив объектов с id, name, settings (JSON), active.  

* __GET /api/v1/lighting-scenarios/{id}__  
Описание: Получить детали сценария.  
Параметры: Path (id сценария).  
Ответ: Объект с полными атрибутами.  

* __POST /api/v1/lighting-scenarios__  
Описание: Создать новый сценарий (например, "Вечерний режим").  
Тело запроса: JSON с lighting_system_id, name, settings (JSON с яркостью и временем).  
Ответ: Созданный объект с id.  

* __PUT /api/v1/lighting-scenarios/{id}__  
Описание: Обновить сценарий.  
Тело запроса: JSON с обновляемыми полями.  
Ответ: Обновленный объект.  

* __DELETE /api/v1/lighting-scenarios/{id}__  
Описание: Удалить сценарий.  
Ответ: 204 No Content.  