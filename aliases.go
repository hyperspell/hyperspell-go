// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"github.com/hyperspell/hyperspell-go/internal/apierror"
	"github.com/hyperspell/hyperspell-go/packages/param"
	"github.com/hyperspell/hyperspell-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type Metadata = shared.Metadata

// This is an alias to an internal type.
type MetadataStatus = shared.MetadataStatus

// Equals "pending"
const MetadataStatusPending = shared.MetadataStatusPending

// Equals "processing"
const MetadataStatusProcessing = shared.MetadataStatusProcessing

// Equals "completed"
const MetadataStatusCompleted = shared.MetadataStatusCompleted

// Equals "failed"
const MetadataStatusFailed = shared.MetadataStatusFailed

// Equals "pending_review"
const MetadataStatusPendingReview = shared.MetadataStatusPendingReview

// Equals "skipped"
const MetadataStatusSkipped = shared.MetadataStatusSkipped

// This is an alias to an internal type.
type Notification = shared.Notification

// This is an alias to an internal type.
type NotificationType = shared.NotificationType

// Equals "error"
const NotificationTypeError = shared.NotificationTypeError

// Equals "warning"
const NotificationTypeWarning = shared.NotificationTypeWarning

// Equals "info"
const NotificationTypeInfo = shared.NotificationTypeInfo

// Equals "success"
const NotificationTypeSuccess = shared.NotificationTypeSuccess

// This is an alias to an internal type.
type QueryResult = shared.QueryResult

// This is an alias to an internal type.
type Resource = shared.Resource

// This is an alias to an internal type.
type ResourceSource = shared.ResourceSource

// Equals "reddit"
const ResourceSourceReddit = shared.ResourceSourceReddit

// Equals "notion"
const ResourceSourceNotion = shared.ResourceSourceNotion

// Equals "slack"
const ResourceSourceSlack = shared.ResourceSourceSlack

// Equals "google_calendar"
const ResourceSourceGoogleCalendar = shared.ResourceSourceGoogleCalendar

// Equals "google_mail"
const ResourceSourceGoogleMail = shared.ResourceSourceGoogleMail

// Equals "box"
const ResourceSourceBox = shared.ResourceSourceBox

// Equals "dropbox"
const ResourceSourceDropbox = shared.ResourceSourceDropbox

// Equals "github"
const ResourceSourceGitHub = shared.ResourceSourceGitHub

// Equals "google_drive"
const ResourceSourceGoogleDrive = shared.ResourceSourceGoogleDrive

// Equals "vault"
const ResourceSourceVault = shared.ResourceSourceVault

// Equals "web_crawler"
const ResourceSourceWebCrawler = shared.ResourceSourceWebCrawler

// Equals "trace"
const ResourceSourceTrace = shared.ResourceSourceTrace

// Equals "microsoft_teams"
const ResourceSourceMicrosoftTeams = shared.ResourceSourceMicrosoftTeams

// Equals "gmail_actions"
const ResourceSourceGmailActions = shared.ResourceSourceGmailActions

// Equals "granola"
const ResourceSourceGranola = shared.ResourceSourceGranola

// Equals "fathom"
const ResourceSourceFathom = shared.ResourceSourceFathom

// Equals "linear"
const ResourceSourceLinear = shared.ResourceSourceLinear
