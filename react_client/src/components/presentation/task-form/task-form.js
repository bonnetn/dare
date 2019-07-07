import React from 'react';

import PropTypes from 'prop-types';

class TaskForm extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            name: this.props.name,
            content: this.props.content,
        };

        this.handleNameChange = this.handleNameChange.bind(this);
        this.handleContentChange = this.handleContentChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }


    handleNameChange(event) {
        this.setState({
            ...this.state,
            name: event.target.value,
        })
    }

    handleContentChange(event) {
        this.setState({
            ...this.state,
            content: event.target.value,
        })
    }

    handleSubmit(event) {
        event.preventDefault();
        this.props.onSubmit(this.state);
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <label>
                    Name:

                    <input type="text" value={this.state.name} onChange={this.handleNameChange}/>
                </label>
                <label>
                    Content:
                    <input type="text" value={this.state.content} onChange={this.handleContentChange}/>
                </label>
                <input type="submit" value="Submit" hidden/>
            </form>
        );
    }
}

TaskForm.propTypes = {
    name: PropTypes.string.isRequired,
    content: PropTypes.string.isRequired,
    onSubmit: PropTypes.func.isRequired,
};

export default TaskForm;
