import useSWR from 'swr'
import styles from '../styles/Home.module.css'

const fetcher = (...args) => fetch(...args).then((res) => res.json())

export default function Currencies() {
    const { data, error } = useSWR(
        ['https://cryptoautotrader-production.up.railway.app/v1/balance', 
        {
        headers: {
           'x-user-id': 'd78964af-9dbe-4613-bf46-f3701bdd0494',
        }
    }], fetcher)

 
    if (error) return <p>could not fetch balance</p>
    if (!data) return <p>no balance data</p>

    return (
        <div className={styles.balance}> 
            <h1> Your Balances </h1>
            <ul>
                {
                    data.balance.map((balance) =>
                    <li key={balance.currency.id}>
                        <div>  {balance.currency.name} balance: {balance.amount}</div>                        
                    </li>
                    )
                }
            </ul>
         </div>
    )
}