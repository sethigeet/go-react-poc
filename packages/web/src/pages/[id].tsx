import Head from "next/head";
import { GetServerSideProps } from "next";
import { User } from "../types/user";
import styles from "./[id].module.scss";

interface Props {
  user: User;
}

export default function Home({ user }: Props) {
  return (
    <>
      <Head>
        <title>User {user.username}</title>
      </Head>
      <div className={styles.container}>
        <h1>User</h1>
        <pre className={styles.userData}>{JSON.stringify(user, null, 2)}</pre>
      </div>
    </>
  );
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const res = await fetch(
    process.env.NEXT_PUBLIC_API_HOST + "/user/" + ctx.params?.id
  );

  const { user } = await res.json();

  return {
    props: {
      user: user ? user : null,
    },
  };
};
