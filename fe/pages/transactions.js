import useSWR from 'swr'
import styles from '../styles/Home.module.css'

const fetcher = (...args) => fetch(...args).then((res) => res.json())

export default function Transactions() {
    const { data, error } = useSWR(
        ['https://cryptoautotrader-production.up.railway.app/v1/transactions?limit=10', 
        {
        headers: {
           'x-user-id': 'd78964af-9dbe-4613-bf46-f3701bdd0494',
        }
    }], fetcher)

 
   if (error) return <p>could not fetch transactions</p>
   if (!data) return <p>no transaction data</p>



    return (
        <div  className={styles.transactions}>
            <h1> Your Transactions </h1>
          <ul>
            {
                data.transactions.map((transaction) => 
                <li className={styles.transactionList} key={transaction.id}>
                    <div> Type: {transaction.transaction_type} </div>
                    <div> Amount: {transaction.amount} </div>
                    <div> Price: {transaction.price} </div>
                </li>
                )
            }
          </ul>
        </div>
    )

}

