import Link from "next/link";

const AuthLayout = ({children}: { children: React.ReactNode }) => {
  return (
    <section>
      <nav><Link href="/auth/login">Login</Link></nav>
      <nav><Link href="/auth/register">Register</Link></nav>
      {children}
    </section>
  );
}

export default AuthLayout;
