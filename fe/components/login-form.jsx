import styles from './login-form.module.css'
import Router from 'next/router'
import { useGlobalContext } from '../contexts';

export default function LoginForm() {
    const { userID, setUserID } = useGlobalContext();
    console.log('login form: current user id ' + userID)

    // Handles the submit event on form submit.
    const handleSubmit = async (event) => {
      // Stop the form from submitting and refreshing the page.
      event.preventDefault()

      // Get data from the form.
      const data = {
        first: event.target.first.value,
        last: event.target.last.value,
      }

      // Send the data to the server in JSON format.
      const JSONdata = JSON.stringify(data)

      // API endpoint where we send form data.
      const endpoint = 'https://cryptoautotrader-production.up.railway.app/v1/users/authenticate'

      // Form the request for sending data to the server.
      const options = {
        // The method is POST because we are sending data.
        method: 'POST',
        // Tell the server we're sending JSON.
        headers: {
          'Content-Type': 'application/json',
        },
        // Body of the request is the JSON data we created above.
        body: JSONdata,
      }

      // Send the form data to our forms API on Vercel and get a response.
      const response = await fetch(endpoint, options)
        .then((response) => {
          if (response.status >= 200 && response.status <= 299) {
            return response.json()
          } else {
            console.log('reject promoise' + response)
            return Promise.reject(response)
          }
        })
        .then((data) => {
            setUserID(data.id)
            console.log('setting user id ', data.id)
            Router.push('/dashboard')
        }).catch((response) => {
            console.log("catch", response.status, response.statusText)
        })
    }


    return (
      <form onSubmit={handleSubmit} method="post">
        <label htmlFor="first">First Name</label>
        <input type="text" id="first" name="first" required />

        <label htmlFor="last">Last Name</label>
        <input type="text" id="last" name="last" required />

        <button type="submit">Submit</button>
      </form>
    )
}
