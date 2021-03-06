/*
 * Minio Client (C) 2014, 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

type errUnexpected struct{}

func (e errUnexpected) Error() string {
	return "Unexpected control flow, please report this bug at https://github.com/minio/mc/issues."
}

type errInvalidSessionID struct {
	id string
}

func (e errInvalidSessionID) Error() string {
	return "Invalid session id ‘" + e.id + "’."
}

type errInvalidACL struct {
	acl string
}

func (e errInvalidACL) Error() string {
	return "Invalid ACL Type ‘" + e.acl + "’."
}

type errNotConfigured struct{}

func (e errNotConfigured) Error() string {
	return "‘mc’ not configured."
}

type errNotAnObject struct {
	url string
}

func (e errNotAnObject) Error() string {
	return "Not an object " + e.url
}

type errInvalidArgument struct{}

func (e errInvalidArgument) Error() string {
	return "Invalid argument."
}

type errUnsupportedScheme struct {
	scheme string
	url    string
}

func (e errUnsupportedScheme) Error() string {
	return "Unsuppported URL scheme: " + e.scheme
}

type errInvalidGlobURL struct {
	glob    string
	request string
}

func (e errInvalidGlobURL) Error() string {
	return "Error reading glob URL " + e.glob + " while comparing with " + e.request
}

type errInvalidAliasName struct {
	name string
}

func (e errInvalidAliasName) Error() string {
	return "Not a valid alias name: " + e.name + " valid examples are: Area51, Grand-Nagus.."
}

type errInvalidAuth struct{}

func (e errInvalidAuth) Error() string {
	return "Invalid auth keys"
}

type errNoMatchingHost struct{}

func (e errNoMatchingHost) Error() string {
	return "No matching host found."
}

type errConfigExists struct{}

func (e errConfigExists) Error() string {
	return "Already exists."
}

// errAliasExists - alias exists
type errAliasExists struct{}

func (e errAliasExists) Error() string {
	return "Already exists."
}

type errInvalidURL struct {
	URL string
}

func (e errInvalidURL) Error() string {
	return "Invalid url " + e.URL
}

type errInvalidSource errInvalidURL

func (e errInvalidSource) Error() string {
	return "Invalid source " + e.URL
}

type errInvalidTarget errInvalidURL

func (e errInvalidTarget) Error() string {
	return "Invalid target " + e.URL
}

type errInvalidTheme struct {
	Theme string
}

func (e errInvalidTheme) Error() string {
	return "Theme " + e.Theme + " is not supported."
}

type errTargetIsNotDir errInvalidURL

func (e errTargetIsNotDir) Error() string {
	return "Target ‘" + e.URL + "’ is not a directory."
}

type errTargetNotFound errInvalidURL

func (e errTargetNotFound) Error() string {
	return "Target directory ‘" + e.URL + "’ does not exist."
}

type errSourceNotRecursive errInvalidURL

func (e errSourceNotRecursive) Error() string {
	return "Source ‘" + e.URL + "’ is not recursive."
}

type errSourceIsNotDir errTargetIsNotDir

func (e errSourceIsNotDir) Error() string {
	return "Source ‘" + e.URL + "’ is not a directory."
}

type errSourceListEmpty errInvalidArgument

func (e errSourceListEmpty) Error() string {
	return "Source list is empty."
}
