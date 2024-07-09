### Endpoints

##### Other
1. `[GET] /health` - get health status

##### Auth
1. `[POST] /auth/login` - logs the user in
2. `[POST] /auth/refresh` - refreshes the access token
3. `[POST] /auth/revoke` - revokes the refresh token
4. `[POST] /auth/logout` - logs the user out
5. `[POST] /auth/link/{provider}` - link the user to the account of the provider

##### Dailies
1. `[GET] /tasks/shared?kind=daily` - get the daily tasks
2. `[GET] /tasks/shared?kind=weekly` - get the weekly tasks
3. `[GET] /tasks/mine?kind=daily` - get the daily tasks for the current user
4. `[GET] /tasks/mine?kind=weekly` - get the weekly tasks for the current user
5. `[PUT] /tasks/shared/{taskID}` - toggle the shared task
6. `[PUT] /tasks/mine/{taskID}` - toggle the task of the current user
7. `[POST] /tasks/mine` - create a new task for the current user

##### Character
1. `[GET] /characters/mine` - get the character for the current user
2. `[POST] /characters/mine` - add the character from the lodestone to the database

##### Settings
1. `[GET] /profile` - get the user's profile
2. `[PUT] /profile` - update the user's profile
