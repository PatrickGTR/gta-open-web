import Head from 'next/head'


export default ({title}) => {
    return (
        <Head>
            <title>{title}</title>
            <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/milligram/1.4.1/milligram.css" />
        </Head>
    )
};
