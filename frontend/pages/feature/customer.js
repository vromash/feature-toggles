import { PureComponent } from 'react';
import axios from 'axios';
import {
    FormControl,
    Input,
    InputLabel,
    Button,
    Typography,
    Paper,
    Grid,
    Checkbox,
} from '@material-ui/core';
import { camelCase, snakeCase } from 'change-case';
import Layout from '../components/Layout';
import styles from '../../styles/FeatureForm.module.scss';

export default class AddCustomerToFeature extends PureComponent {
    constructor(props) {
        super(props);

        this.state = {
            customerId: 0,
            featureId: 0,
            isActive: false,
        };
    }

    handleSubmit = async () => {
        await axios.post('/api/add-customer-to-feature', {
            customerId: parseInt(this.state.customerId),
            featureId: parseInt(this.state.featureId),
            active: this.state.isActive,
        });
    }

    handleStateChange = (event) => {
        this.setState(() => ({
            [camelCase(event.target.id)]: event.target.value,
        }));
    };

    handleCheckboxChange = (event) => {
        this.setState(() => ({
            [camelCase(event.target.id)]: event.target.checked,
        }));
    }

    renderNumberInput(text, size) {
        return (
            <Grid item xs={size}>
                <FormControl>
                    <InputLabel htmlFor={snakeCase(text)}>{text}</InputLabel>
                    <Input id={snakeCase(text)} onChange={this.handleStateChange} type="number"/>
                </FormControl>
            </Grid>
        );
    }

    renderCheckbox(text) {
        return (
            <Grid
                item
                xs={6}
                style={{ display: 'flex' }}
            >
                <Typography>{text}</Typography>
                <Checkbox
                    checked={this.state[camelCase(text)]}
                    onChange={this.handleCheckboxChange}
                    id={snakeCase(text)}
                />
            </Grid>
        );
    }

    render() {
        return (
            <Layout>
                <div>
                    <Paper className={styles.paper}>
                        <Typography variant="h4">Add customer to feature</Typography>
                        <Grid
                            alignItems="center"
                            alignContent="center"
                            justifyContent="flex-start"
                            container
                            spacing={1}
                        >
                            {this.renderNumberInput('Customer ID', 6)}
                            {this.renderNumberInput('Feature ID', 6)}
                            {this.renderCheckbox('Is active')}
                            <Grid item xs={12}>
                                <Button variant="contained" color="secondary" onClick={this.handleSubmit}>Submit</Button>
                            </Grid>
                        </Grid>
                    </Paper>
                </div>
            </Layout>
        );
    }
}
