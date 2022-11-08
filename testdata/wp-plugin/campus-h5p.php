<?php
/**
 * Campus H5P main plugin file
 *
 * @package campus-h5p
 */

/**
 * Plugin Name: H5P Campus
 * Description: H5P Campus plugin
 * Version: 1.0.0-beta.5
 * Text Domain: h5p-campus
 * Author: Campus team
 */

namespace Campus\H5P;

define( 'CAMPUS_H5P_VERSION', '1.0.0-beta.5' );

define( __NAMESPACE__ . '\PLUGIN_FILE', __FILE__ );
define( __NAMESPACE__ . '\PLUGIN_DIR', basename( dirname( __FILE__ ) ) );
define( __NAMESPACE__ . '\PLUGIN_FILENAME', basename( __FILE__ ) );

require_once dirname( __FILE__ ) . '/src/exceptions.php';
require_once dirname( __FILE__ ) . '/src/loader.php';

Main::get()->boot();
