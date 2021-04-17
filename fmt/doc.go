/*
* Package fmt is 100% compatible with fmt but changes strings to
* interface{} and allows a 'func() string' to be passed as well. This is
* useful at the highest level for generating help information and other
* text that can contain enclosed and package global data instead of
* a static string.
*
* Also, nil values of any kind will be translated to empty strings ("").
*
* This package is guaranteed to remain 100% compatible
* as long as the standard fmt package allows.
 */
package fmt
