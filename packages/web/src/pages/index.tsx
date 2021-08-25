import Head from "next/head";
import { GetServerSideProps } from "next";
import { useRouter } from "next/router";

import { User } from "../types/user";
import styles from "@/styles/pages/index.module.scss";

interface Props {
  users: User[];
}

export default function Home({ users }: Props) {
  const router = useRouter();

  return (
    <>
      <Head>
        <title>Users</title>
      </Head>
      <div className={styles.container}>
        <h1>Users</h1>
        <table className={styles.usersTable}>
          <thead>
            <tr>
              <th>Username</th>
              <th>Email</th>
              <th>Created At</th>
              <th>Updated At</th>
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
    </>
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
