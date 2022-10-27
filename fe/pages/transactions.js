import { useState, useEffect } from 'react'

export default function Transactions() {

    const [data, setData] = useState(null)
    const [isLoading, setLoading] = useState(false)

    useEffect(() => {
        setLoading(true)
        fetch('https://cryptoautotrader-production.up.railway.app/v1/transactions', {
            headers: {
               'x-user-id': 'd78964af-9dbe-4613-bf46-f3701bdd0494',
            }
        })
            .then((res) => {
                if (res.ok) {
                    return res.json()
                }
                throw new Error('failed to retrieve data ')
            })
            .then((data) => {
                setData(data)
                setLoading(false)
            })
            .catch((error) => {
                console.log('network failure ' + error)
            })
    }, [])
     
  
   if (isLoading) return <p>Loading...</p>
   if (!data) return <p>No profile data</p>

    return (
        <div>
            <h1> Your Transactions </h1>
          <ul>
            {
                data.transactions.map((transaction) => (
                    <li>{transaction.id}</li>
                    )
                )
            }
          </ul>
        </div>
    )
}

