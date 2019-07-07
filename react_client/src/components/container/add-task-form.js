import React from 'react';

import {connect} from "react-redux";
import TaskForm from "../presentation/task-form/task-form";
import {addTask} from "../../store/actions";

const AddTaskFormBase = ({addTask}) => {
    return <TaskForm name='lol' content='test' onSubmit={addTask}/>
};

const mapDispatchToProps = (dispatch) => {
    return {
        addTask: ({name, content}) => {
            dispatch(addTask(name, content))
        },
    }
};

const AddTaskForm = connect(
    null,
    mapDispatchToProps,
)(AddTaskFormBase);

export default AddTaskForm;