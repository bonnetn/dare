import React from 'react';

import {connect} from "react-redux";
import {deleteTask} from "../../store/actions";

import PropTypes from 'prop-types';

class DeleteTaskFormBase extends React.Component {
    constructor(props) {
        super(props);
        this.onDelete = this.onDelete.bind(this);
    }

    onDelete(event) {
        if (!confirm("Are you sure?")) { //eslint-disable-line
            return
        }
        this.props.deleteTask(this.props.uuid);
        event.preventDefault();
    }

    render() {
        return <button onClick={this.onDelete}>DELETE</button>
    }
}

DeleteTaskFormBase.propTypes = {
    deleteTask: PropTypes.func.isRequired,
    uuid: PropTypes.string.isRequired,
};

const mapDispatchToProps = (dispatch) => {
    return {
        deleteTask: (uuid) => {
            dispatch(deleteTask(uuid))
        },
    }
};

const DeleteTaskForm = connect(
    null,
    mapDispatchToProps,
)(DeleteTaskFormBase);

export default DeleteTaskForm;