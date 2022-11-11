import Head from 'next/head'
import 'bootstrap/dist/css/bootstrap.css'
import LoginForm from '../components/login-form'

export default function Login() {
    return (
        <div>
        <Head>
          <title>Login Page</title>
          <link rel="icon" href="/favicon.ico" />
        </Head>
        <header>
            <h1> Welcome to crypto auto bot </h1>
        </header>
        <main>
            <h5> please login to continue </h5>
            <LoginForm/>
        </main>
      </div>
    )
}