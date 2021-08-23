import Head from 'next/head';
import { PureComponent } from 'react';
import styles from '../../styles/Layout.module.scss';

class Layout extends PureComponent {
    renderHead() {
        return (
            <Head>
                <title>Feature Toggles</title>
            </Head>
        );
    }

    render() {
        const { children } = this.props;
        return (
            <main className={styles.main}>
                <div className={styles.container}>
                    {this.renderHead()}
                    {children}
                </div>
            </main>
        );
    }
}

export default Layout;
