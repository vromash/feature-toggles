import { PureComponent } from 'react';
import axios from 'axios';
import FeatureForm from '../components/FeatureForm';
import Layout from '../components/Layout';

export default class EditFeature extends PureComponent {
    async onSubmit(data) {
        const { id, ...rest } = data;
        await axios.post(
            '/api/edit-feature',
            {
                id: parseInt(id),
                ...rest,
            },
        );
    }

    render() {
        return (
            <Layout>
                <FeatureForm
                    buttonText="Edit"
                    onSubmit={this.onSubmit}
                    title="Edit feature"
                    includeId={true}
                />
            </Layout>
        );
    }
}
