import useSWR from 'swr'

const fetcher = (...args) => fetch(...args).then((res) => res.json())

export default function Transactions() {
    const { data, error } = useSWR(
        ['https://cryptoautotrader-production.up.railway.app/v1/transactions', 
        {
        headers: {
           'x-user-id': 'd78964af-9dbe-4613-bf46-f3701bdd0494',
        }
    }], fetcher)

 
   if (error) return <p>could not fetch data</p>
   if (!data) return <p>No profile data</p>

    return (
        <div>
            <h1> Your Transactions </h1>
          <ul>
            {
                data.transactions.map((transaction) => 
                <li key={transaction.id}>
                        Type: {transaction.transaction_type} 
                        Amount: {transaction.amount} 
                        Price: {transaction.price} 
                </li>
                )
              
            }
          </ul>
        </div>
    )
}

