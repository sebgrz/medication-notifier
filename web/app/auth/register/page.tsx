'use client';

import { useApiManager } from "@/hooks/useApiManager";
import { Field, Form, Formik } from "formik";
import { useRouter } from "next/navigation";

const Register = () => {
  const api = useApiManager();
  const router = useRouter();
  return (
    <div>
      <Formik
        initialValues={{ username: '', password: '' }}
        onSubmit={async (values, { setSubmitting }) => {
          setSubmitting(true);
          if (await api.authRegister(values.username, values.password)) {
            router.push("/auth/login");
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

export default Register;
