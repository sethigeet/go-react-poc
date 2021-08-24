import { User } from "../types/user";
import { useRouter } from "next/router";
import styles from "./index.module.scss";
import { GetServerSideProps } from "next";

interface Props {
  users: User[];
}

export default function Home({ users }: Props) {
  const router = useRouter();

  return (
    <div className={styles.container}>
      <h1>Users</h1>
      <table className={styles.usersTable}>
        <thead>
          <tr>
            <th>Username</th>
            <th>Email</th>
            <th>CreatedAt</th>
            <th>UpdatedAt</th>
          </tr>
        </thead>
        <tbody>
          {users.map((user) => (
            <tr key={user.id} onClick={() => router.push("/" + user.id)}>
              <td>{user.username}</td>
              <td>{user.email}</td>
              <td>{new Date(user.createdAt).toLocaleString()}</td>
              <td>{new Date(user.updatedAt).toLocaleString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export const getServerSideProps: GetServerSideProps = async () => {
  const res = await fetch(process.env.NEXT_PUBLIC_API_HOST + "/user/");
  const { users } = await res.json();

  return {
    props: {
      users,
    },
  };
};
