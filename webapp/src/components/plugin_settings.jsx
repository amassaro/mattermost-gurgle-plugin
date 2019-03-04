import PropTypes from 'prop-types';
import React from 'react';

export default class PluginSettings extends React.Component {
    static propTypes = {
        teamId: PropTypes.node.isRequired,
        channelName: PropTypes.node.isRequired,
    };

    static defaultProps = {
    };

    shouldComponentUpdate() {
        return true;
    }

    render() {
        return (
            <span>{this.props.teamId} {this.props.channelName}</span>
        );
    }
}
