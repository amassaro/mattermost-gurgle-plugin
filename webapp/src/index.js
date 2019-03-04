import {id as pluginId} from './manifest';

//import PlugingSettings from './components/plugin_settings';

export default class GurglePlugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {
        // @see https://developers.mattermost.com/extend/plugins/webapp/reference
        //registry.registerPluginSettingsComponent(PlugingSettings);
    }
}

window.registerPlugin(pluginId, new GurglePlugin());
