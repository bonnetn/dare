import React from 'react';

import PropTypes from 'prop-types';

class TaskForm extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            name: props.task.name,
        };

        this.handleNameChange = this.handleNameChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }


    handleNameChange(event) {
        this.setState({
            ...this.state,
            name: event.target.value,
        })
    }

    handleSubmit(event) {
        event.preventDefault();
        this.props.onSubmit({
            ...this.props.task,
            name: this.state.name,
        });
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <label>
                    Name:

                    <input type="text" value={this.state.name} onChange={this.handleNameChange}/>
                </label>
                <input type="submit" value="Submit" hidden/>
            </form>
        );
    }
}

TaskForm.propTypes = {
    task: PropTypes.object.isRequired,
    onSubmit: PropTypes.func.isRequired,
};

export default TaskForm;
