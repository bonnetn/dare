import React from 'react';

import {connect} from "react-redux";
import TaskForm from "../presentation/task-form/task-form";
import {editTask} from "../../store/actions";

import PropTypes from 'prop-types';

class EditTaskFormBase extends React.Component {
    constructor(props) {
        super(props);
        this.onSubmit = this.onSubmit.bind(this);
    }
    onSubmit({name, content}) {
        const task = this.props.task;
        this.props.editTask(task.uuid, task.version, name, content);
    }

    render() {
        return <TaskForm name={this.props.task.name} content={this.props.task.content} onSubmit={this.onSubmit}/>
    }
}

EditTaskFormBase.propTypes = {
    editTask: PropTypes.func.isRequired,
    task: PropTypes.shape({
        uuid: PropTypes.string.isRequired,
        version: PropTypes.number.isRequired,
        name: PropTypes.string.isRequired,
        content: PropTypes.string.isRequired,
    }).isRequired,
};

const mapDispatchToProps = (dispatch) => {
    return {
        editTask: (uuid, version, name, content) => {
            dispatch(editTask(uuid, version, name, content))
        },
    }
};

const EditTaskForm = connect(
    null,
    mapDispatchToProps,
)(EditTaskFormBase);

export default EditTaskForm;