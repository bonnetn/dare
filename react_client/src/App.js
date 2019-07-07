import React from 'react';
import './App.css';
import ActiveTaskViewer from "./components/container/active-task-viewer";

class App extends React.Component {
    render() {
        return (
            <div>
                <ActiveTaskViewer/>
            </div>
        );
    }
}

export default App;
