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

    onSubmit(task) {
        this.props.editTask(task);
    }

    render() {
        return <TaskForm task={this.props.task} onSubmit={this.onSubmit}/>
    }
}

EditTaskFormBase.propTypes = {
    editTask: PropTypes.func.isRequired,
    task: PropTypes.shape({
        uuid: PropTypes.string.isRequired,
        name: PropTypes.string.isRequired,
    }).isRequired,
};

const mapDispatchToProps = (dispatch) => {
    return {
        editTask: (task) => {
            dispatch(editTask(task))
        },
    }
};

const EditTaskForm = connect(
    null,
    mapDispatchToProps,
)(EditTaskFormBase);

export default EditTaskForm;