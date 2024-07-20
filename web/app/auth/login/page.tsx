'use client';

import { useApiManager } from "@/hooks/useApiManager";
import { Field, Form, Formik } from "formik";
import { useRouter } from "next/navigation";

const Login = () => {
  const api = useApiManager();
  const router = useRouter();

  return (
    <div>
      Login:<br/>
      <Formik
        initialValues={{ username: '', password: '' }}
        onSubmit={async (values, { setSubmitting }) => {
          setSubmitting(true);
          if (await api.authLogin(values.username, values.password)) {
            router.push("/");
          }
          setSubmitting(false);
        }}
      >
        {({ isSubmitting }) => (
          <Form>
            <Field type="text" name="username" />
            <Field type="password" name="password" />
            <button type="submit" disabled={isSubmitting}>
              Register
            </button>
          </Form>
        )}
      </Formik>
    </div>
  );
}

export default Login;
