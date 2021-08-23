import { PureComponent } from 'react';
import {
    FormControl,
    Input,
    InputLabel,
    Button,
    Typography,
    Checkbox,
    Paper,
    Grid,
} from '@material-ui/core';
import { camelCase, snakeCase } from 'change-case';
import {
    KeyboardTimePicker,
    MuiPickersUtilsProvider,
    KeyboardDatePicker,
} from '@material-ui/pickers';
import DateFnsUtils from '@date-io/date-fns';
import styles from '../../styles/FeatureForm.module.scss';

export default class FeatureForm extends PureComponent {
    constructor(props) {
        super(props);

        this.state = {
            id: 0,
            displayName: '',
            technicalName: '',
            expiresOn: new Date(),
            description: '',
            inverted: false,
        };
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

    handleDateChange = (date) => {
        this.setState(() => ({
            expiresOn: date,
        }));
    }

    renderTextInput(text, size, type = 'text') {
        return (
            <Grid item xs={size}>
                <FormControl>
                    <InputLabel htmlFor={snakeCase(text)}>{text}</InputLabel>
                    <Input id={snakeCase(text)} onChange={this.handleStateChange} type={type}/>
                </FormControl>
            </Grid>
        );
    }

    renderDateTimeInput(text) {
        return (
            <>
                <Grid item xs={6}>
                    <MuiPickersUtilsProvider utils={DateFnsUtils}>
                        <KeyboardDatePicker
                            disableToolbar
                            variant="inline"
                            format="MM/dd/yyyy"
                            margin="normal"
                            id={`${ snakeCase(text) }_date`}
                            label={`${ text }: date`}
                            value={this.state.expiresOn}
                            onChange={this.handleDateChange}
                        />
                    </MuiPickersUtilsProvider>
                </Grid>
                <Grid item xs={6}>
                    <MuiPickersUtilsProvider utils={DateFnsUtils}>
                        <KeyboardTimePicker
                            margin="normal"
                            id={`${ snakeCase(text) }_time`}
                            label={`${ text } time`}
                            value={this.state.expiresOn}
                            onChange={this.handleDateChange}
                        />
                    </MuiPickersUtilsProvider>
                </Grid>
            </>
        );
    }

    renderCheckbox(text) {
        return (
            <Grid
                item
                xs={6}
                alignItems="center"
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

    handleSubmit = () => {
        this.props.onSubmit({
            ...this.state,
        });
    }

    render() {
        const { buttonText, title, includeId } = this.props;

        return (
            <div>
                <Paper className={styles.paper}>
                    <Typography variant="h4">{title}</Typography>
                    <Grid
                        alignContent="center"
                        justifyContent="flex-start"
                        container
                        spacing={1}
                    >
                        {includeId ? this.renderTextInput('Id', 12, 'number') : ''}
                        {this.renderTextInput('Display name', 6)}
                        {this.renderTextInput('Technical name', 6)}
                        {this.renderDateTimeInput('Expires on')}
                        {this.renderTextInput('Description', 6)}
                        {this.renderCheckbox('Inverted')}

                        <Grid item xs={12}>
                            <Button variant="contained" color="secondary" onClick={this.handleSubmit}>{buttonText}</Button>
                        </Grid>
                    </Grid>
                </Paper>
            </div>
        );
    }
}
