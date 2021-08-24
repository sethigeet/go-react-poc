import Link from "next/link";
import { useRouter } from "next/router";
import styles from "./Navbar.module.scss";

export const Navbar: React.FC<{}> = () => {
  const { pathname } = useRouter();

  return (
    <div className={styles.navbar}>
      <h1>Go React POC</h1>
      <ul>
        {[
          { name: "Home", link: "/" },
          { name: "Create New", link: "/create" },
        ].map(({ name, link }, i) => (
          <li
            key={i}
            style={{ fontWeight: pathname === link ? "bold" : undefined }}
          >
            <Link href={link}>{name}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
};
