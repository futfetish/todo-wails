import { FC, ReactNode } from "react";
import { NavLink } from "react-router-dom";
import Styles from "../styles/nav.module.scss";
import clsx from "clsx";

export const Nav: FC = () => {
  return (
    <nav className={Styles.nav}>
      <But to={"/"}>home</But>
      <But to={"/completed"}>completed</But>
    </nav>
  );
};

const But: FC<{ children: ReactNode } & React.ComponentProps<typeof NavLink>> = ({
  children,
  ...props
}) => {
  return (
    <NavLink className={ ({isActive}) => clsx( Styles.but , isActive && Styles.butActive )} {...props}>
      {children}
    </NavLink>
  );
};
