### Endpoints

##### Other
- [x] `[GET] /health` - get health status

##### Auth
- [x] `[POST] /auth/login` - logs the user in
- [x] `[POST] /auth/refresh` - refreshes the access token
- [x] `[POST] /auth/revoke` - revokes the refresh token
- [ ] `[POST] /auth/logout` - logs the user out
- [ ] `[POST] /auth/link/{provider}` - link the user to the account of the provider

##### Dailies
- [x] `[GET] /tasks/shared?kind=daily` - get the daily tasks
- [x] `[GET] /tasks/shared?kind=weekly` - get the weekly tasks
- [ ] `[GET] /tasks/mine?kind=daily` - get the daily tasks for the current user
- [ ] `[GET] /tasks/mine?kind=weekly` - get the weekly tasks for the current user
- [x] `[PUT] /tasks/shared/{taskID}` - toggle the shared task
- [ ] `[PUT] /tasks/mine/{taskID}` - toggle the task of the current user
- [ ] `[POST] /tasks/mine` - create a new task for the current user

##### Character
- [ ] `[GET] /characters/mine` - get the character for the current user
- [ ] `[POST] /characters/mine` - add the character from the lodestone to the database

##### Settings
- [ ] `[GET] /profile` - get the user's profile
- [ ] `[PUT] /profile` - update the user's profile

##### Misc
- [x] `[GET] /events` - endpoint for SSE events
