import { FormEventHandler, useState } from "react";
import Head from "next/head";
import { useRouter } from "next/router";

import styles from "@/styles/pages/create.module.scss";

export default function Home() {
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");

  const [errors, setErrors] = useState<Record<string, string>>({});

  const router = useRouter();

  const onSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();

    const res = await fetch(process.env.NEXT_PUBLIC_API_HOST + "/user/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, email }),
    });
    const data = await res.json();

    // Check for errors
    if (data.errors) {
      setErrors(data.errors);
      return;
    }

    // Send the user to the home page
    router.push("/");
  };

  return (
    <>
      <Head>
        <title>Create a New User</title>
      </Head>
      <div className={styles.container}>
        <h1>Create a New User</h1>

        <form onSubmit={onSubmit} className={styles.createForm}>
          <div>
            <label htmlFor="username">Username</label>
            <input
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              id="username"
            />
            <span className="error">{errors.username || ""}&nbsp;</span>
          </div>
          <div>
            <label htmlFor="email">Email</label>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              id="email"
            />
            <span className="error">{errors.email || ""}&nbsp;</span>
          </div>

          <button type="submit" className={styles.createBtn}>
            Create
          </button>
        </form>
      </div>
    </>
  );
}
