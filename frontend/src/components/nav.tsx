import { FC, ReactNode } from "react";
import { Link } from "react-router-dom";
import Styles from "../styles/nav.module.scss";

export const Nav: FC = () => {
  return (
    <nav className={Styles.nav}>
      <But to={"/"}>home</But>
      <But to={"/completed"}>completed</But>
    </nav>
  );
};

const But: FC<{ children: ReactNode } & React.ComponentProps<typeof Link>> = ({
  children,
  ...props
}) => {
  return (
    <Link className={Styles.but} {...props}>
      {children}
    </Link>
  );
};
