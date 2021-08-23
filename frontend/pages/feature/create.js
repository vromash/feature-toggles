import { PureComponent } from 'react';
import axios from 'axios';
import FeatureForm from '../components/FeatureForm';
import Layout from '../components/Layout';

export default class CreateFeature extends PureComponent {
    async onSubmit(data) {
        const { id, ...rest } = data;
        await axios.post('/api/create-feature', rest);
    }

    render() {
        return (
            <Layout>
                <FeatureForm
                    buttonText="Create"
                    onSubmit={this.onSubmit}
                    title="Create new feature"
                />
            </Layout>
        );
    }
}
