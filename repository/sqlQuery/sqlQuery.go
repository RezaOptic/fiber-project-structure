package sqlQuery

const InsertNewUserQuery = `insert into public.user(username, created_at, updated_at) VALUES ($1,$2,$3) RETURNING id;`
const GetUsersQuery = `select id, username, created_at, updated_at from public.user order by updated_at %s limit %d offset %d;`
const GetUserQuery = `select id, username, created_at, updated_at from public.user where id=$1 and deleted=false;`
const UpdateUserQuery = `update public.user set username=$1, updated_at=$2 where id=$3 and deleted=false;`
const DeleteUserQuery = `update public.user set deleted=true, updated_at=now() where id=$1;`
