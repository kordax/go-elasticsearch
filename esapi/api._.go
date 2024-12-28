// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// Code generated from specification version 8.16.0 (30399bb): DO NOT EDIT

package esapi

// API contains the Elasticsearch APIs
type API struct {
	Remote *Remote

	Bulk                                          Bulk
	Capabilities                                  Capabilities
	ClearScroll                                   ClearScroll
	ConnectorCheckIn                              ConnectorCheckIn
	ConnectorDelete                               ConnectorDelete
	ConnectorGet                                  ConnectorGet
	ConnectorLastSync                             ConnectorLastSync
	ConnectorList                                 ConnectorList
	ConnectorPost                                 ConnectorPost
	ConnectorPut                                  ConnectorPut
	ConnectorSecretDelete                         ConnectorSecretDelete
	ConnectorSecretGet                            ConnectorSecretGet
	ConnectorSecretPost                           ConnectorSecretPost
	ConnectorSecretPut                            ConnectorSecretPut
	ConnectorSyncJobCancel                        ConnectorSyncJobCancel
	ConnectorSyncJobCheckIn                       ConnectorSyncJobCheckIn
	ConnectorSyncJobClaim                         ConnectorSyncJobClaim
	ConnectorSyncJobDelete                        ConnectorSyncJobDelete
	ConnectorSyncJobError                         ConnectorSyncJobError
	ConnectorSyncJobGet                           ConnectorSyncJobGet
	ConnectorSyncJobList                          ConnectorSyncJobList
	ConnectorSyncJobPost                          ConnectorSyncJobPost
	ConnectorSyncJobUpdateStats                   ConnectorSyncJobUpdateStats
	ConnectorUpdateAPIKeyDocumentID               ConnectorUpdateAPIKeyDocumentID
	ConnectorUpdateActiveFiltering                ConnectorUpdateActiveFiltering
	ConnectorUpdateConfiguration                  ConnectorUpdateConfiguration
	ConnectorUpdateError                          ConnectorUpdateError
	ConnectorUpdateFeatures                       ConnectorUpdateFeatures
	ConnectorUpdateFiltering                      ConnectorUpdateFiltering
	ConnectorUpdateFilteringValidation            ConnectorUpdateFilteringValidation
	ConnectorUpdateIndexName                      ConnectorUpdateIndexName
	ConnectorUpdateName                           ConnectorUpdateName
	ConnectorUpdateNative                         ConnectorUpdateNative
	ConnectorUpdatePipeline                       ConnectorUpdatePipeline
	ConnectorUpdateScheduling                     ConnectorUpdateScheduling
	ConnectorUpdateServiceDocumentType            ConnectorUpdateServiceDocumentType
	ConnectorUpdateStatus                         ConnectorUpdateStatus
	Count                                         Count
	Create                                        Create
	DanglingIndicesDeleteDanglingIndex            DanglingIndicesDeleteDanglingIndex
	DanglingIndicesImportDanglingIndex            DanglingIndicesImportDanglingIndex
	DanglingIndicesListDanglingIndices            DanglingIndicesListDanglingIndices
	DeleteByQuery                                 DeleteByQuery
	DeleteByQueryRethrottle                       DeleteByQueryRethrottle
	Delete                                        Delete
	DeleteScript                                  DeleteScript
	Exists                                        Exists
	ExistsSource                                  ExistsSource
	Explain                                       Explain
	FeaturesGetFeatures                           FeaturesGetFeatures
	FeaturesResetFeatures                         FeaturesResetFeatures
	FieldCaps                                     FieldCaps
	FleetDeleteSecret                             FleetDeleteSecret
	FleetGetSecret                                FleetGetSecret
	FleetGlobalCheckpoints                        FleetGlobalCheckpoints
	FleetMsearch                                  FleetMsearch
	FleetPostSecret                               FleetPostSecret
	FleetSearch                                   FleetSearch
	Get                                           Get
	GetScriptContext                              GetScriptContext
	GetScriptLanguages                            GetScriptLanguages
	GetScript                                     GetScript
	GetSource                                     GetSource
	HealthReport                                  HealthReport
	Index                                         Index
	InferenceDelete                               InferenceDelete
	InferenceGet                                  InferenceGet
	InferenceInference                            InferenceInference
	InferencePut                                  InferencePut
	InferenceStreamInference                      InferenceStreamInference
	Info                                          Info
	KnnSearch                                     KnnSearch
	Mget                                          Mget
	Msearch                                       Msearch
	MsearchTemplate                               MsearchTemplate
	Mtermvectors                                  Mtermvectors
	Ping                                          Ping
	ProfilingStacktraces                          ProfilingStacktraces
	ProfilingStatus                               ProfilingStatus
	ProfilingTopnFunctions                        ProfilingTopnFunctions
	PutScript                                     PutScript
	QueryRulesDeleteRule                          QueryRulesDeleteRule
	QueryRulesDeleteRuleset                       QueryRulesDeleteRuleset
	QueryRulesGetRule                             QueryRulesGetRule
	QueryRulesGetRuleset                          QueryRulesGetRuleset
	QueryRulesListRulesets                        QueryRulesListRulesets
	QueryRulesPutRule                             QueryRulesPutRule
	QueryRulesPutRuleset                          QueryRulesPutRuleset
	QueryRulesTest                                QueryRulesTest
	RankEval                                      RankEval
	Reindex                                       Reindex
	ReindexRethrottle                             ReindexRethrottle
	RenderSearchTemplate                          RenderSearchTemplate
	ScriptsPainlessExecute                        ScriptsPainlessExecute
	Scroll                                        Scroll
	SearchApplicationDeleteBehavioralAnalytics    SearchApplicationDeleteBehavioralAnalytics
	SearchApplicationDelete                       SearchApplicationDelete
	SearchApplicationGetBehavioralAnalytics       SearchApplicationGetBehavioralAnalytics
	SearchApplicationGet                          SearchApplicationGet
	SearchApplicationList                         SearchApplicationList
	SearchApplicationPostBehavioralAnalyticsEvent SearchApplicationPostBehavioralAnalyticsEvent
	SearchApplicationPutBehavioralAnalytics       SearchApplicationPutBehavioralAnalytics
	SearchApplicationPut                          SearchApplicationPut
	SearchApplicationRenderQuery                  SearchApplicationRenderQuery
	SearchApplicationSearch                       SearchApplicationSearch
	SearchMvt                                     SearchMvt
	Search                                        Search
	SearchShards                                  SearchShards
	SearchTemplate                                SearchTemplate
	ShutdownDeleteNode                            ShutdownDeleteNode
	ShutdownGetNode                               ShutdownGetNode
	ShutdownPutNode                               ShutdownPutNode
	SimulateIngest                                SimulateIngest
	SynonymsDeleteSynonym                         SynonymsDeleteSynonym
	SynonymsDeleteSynonymRule                     SynonymsDeleteSynonymRule
	SynonymsGetSynonym                            SynonymsGetSynonym
	SynonymsGetSynonymRule                        SynonymsGetSynonymRule
	SynonymsGetSynonymsSets                       SynonymsGetSynonymsSets
	SynonymsPutSynonym                            SynonymsPutSynonym
	SynonymsPutSynonymRule                        SynonymsPutSynonymRule
	TermsEnum                                     TermsEnum
	Termvectors                                   Termvectors
	UpdateByQuery                                 UpdateByQuery
	UpdateByQueryRethrottle                       UpdateByQueryRethrottle
	Update                                        Update
}

// Remote contains the Remote APIs
type Remote struct {
}

// New creates new API
func New(t Transport) *API {
	return &API{
		Bulk:                               newBulkFunc(t),
		Capabilities:                       newCapabilitiesFunc(t),
		ClearScroll:                        newClearScrollFunc(t),
		ConnectorCheckIn:                   newConnectorCheckInFunc(t),
		ConnectorDelete:                    newConnectorDeleteFunc(t),
		ConnectorGet:                       newConnectorGetFunc(t),
		ConnectorLastSync:                  newConnectorLastSyncFunc(t),
		ConnectorList:                      newConnectorListFunc(t),
		ConnectorPost:                      newConnectorPostFunc(t),
		ConnectorPut:                       newConnectorPutFunc(t),
		ConnectorSecretDelete:              newConnectorSecretDeleteFunc(t),
		ConnectorSecretGet:                 newConnectorSecretGetFunc(t),
		ConnectorSecretPost:                newConnectorSecretPostFunc(t),
		ConnectorSecretPut:                 newConnectorSecretPutFunc(t),
		ConnectorSyncJobCancel:             newConnectorSyncJobCancelFunc(t),
		ConnectorSyncJobCheckIn:            newConnectorSyncJobCheckInFunc(t),
		ConnectorSyncJobClaim:              newConnectorSyncJobClaimFunc(t),
		ConnectorSyncJobDelete:             newConnectorSyncJobDeleteFunc(t),
		ConnectorSyncJobError:              newConnectorSyncJobErrorFunc(t),
		ConnectorSyncJobGet:                newConnectorSyncJobGetFunc(t),
		ConnectorSyncJobList:               newConnectorSyncJobListFunc(t),
		ConnectorSyncJobPost:               newConnectorSyncJobPostFunc(t),
		ConnectorSyncJobUpdateStats:        newConnectorSyncJobUpdateStatsFunc(t),
		ConnectorUpdateAPIKeyDocumentID:    newConnectorUpdateAPIKeyDocumentIDFunc(t),
		ConnectorUpdateActiveFiltering:     newConnectorUpdateActiveFilteringFunc(t),
		ConnectorUpdateConfiguration:       newConnectorUpdateConfigurationFunc(t),
		ConnectorUpdateError:               newConnectorUpdateErrorFunc(t),
		ConnectorUpdateFeatures:            newConnectorUpdateFeaturesFunc(t),
		ConnectorUpdateFiltering:           newConnectorUpdateFilteringFunc(t),
		ConnectorUpdateFilteringValidation: newConnectorUpdateFilteringValidationFunc(t),
		ConnectorUpdateIndexName:           newConnectorUpdateIndexNameFunc(t),
		ConnectorUpdateName:                newConnectorUpdateNameFunc(t),
		ConnectorUpdateNative:              newConnectorUpdateNativeFunc(t),
		ConnectorUpdatePipeline:            newConnectorUpdatePipelineFunc(t),
		ConnectorUpdateScheduling:          newConnectorUpdateSchedulingFunc(t),
		ConnectorUpdateServiceDocumentType: newConnectorUpdateServiceDocumentTypeFunc(t),
		ConnectorUpdateStatus:              newConnectorUpdateStatusFunc(t),
		Count:                              newCountFunc(t),
		Create:                             newCreateFunc(t),
		DanglingIndicesDeleteDanglingIndex: newDanglingIndicesDeleteDanglingIndexFunc(t),
		DanglingIndicesImportDanglingIndex: newDanglingIndicesImportDanglingIndexFunc(t),
		DanglingIndicesListDanglingIndices: newDanglingIndicesListDanglingIndicesFunc(t),
		DeleteByQuery:                      newDeleteByQueryFunc(t),
		DeleteByQueryRethrottle:            newDeleteByQueryRethrottleFunc(t),
		Delete:                             newDeleteFunc(t),
		DeleteScript:                       newDeleteScriptFunc(t),
		Exists:                             newExistsFunc(t),
		ExistsSource:                       newExistsSourceFunc(t),
		Explain:                            newExplainFunc(t),
		FeaturesGetFeatures:                newFeaturesGetFeaturesFunc(t),
		FeaturesResetFeatures:              newFeaturesResetFeaturesFunc(t),
		FieldCaps:                          newFieldCapsFunc(t),
		FleetDeleteSecret:                  newFleetDeleteSecretFunc(t),
		FleetGetSecret:                     newFleetGetSecretFunc(t),
		FleetGlobalCheckpoints:             newFleetGlobalCheckpointsFunc(t),
		FleetMsearch:                       newFleetMsearchFunc(t),
		FleetPostSecret:                    newFleetPostSecretFunc(t),
		FleetSearch:                        newFleetSearchFunc(t),
		Get:                                newGetFunc(t),
		GetScriptContext:                   newGetScriptContextFunc(t),
		GetScriptLanguages:                 newGetScriptLanguagesFunc(t),
		GetScript:                          newGetScriptFunc(t),
		GetSource:                          newGetSourceFunc(t),
		HealthReport:                       newHealthReportFunc(t),
		Index:                              newIndexFunc(t),
		InferenceDelete:                    newInferenceDeleteFunc(t),
		InferenceGet:                       newInferenceGetFunc(t),
		InferenceInference:                 newInferenceInferenceFunc(t),
		InferencePut:                       newInferencePutFunc(t),
		InferenceStreamInference:           newInferenceStreamInferenceFunc(t),
		Info:                               newInfoFunc(t),
		KnnSearch:                          newKnnSearchFunc(t),
		Mget:                               newMgetFunc(t),
		Msearch:                            newMsearchFunc(t),
		MsearchTemplate:                    newMsearchTemplateFunc(t),
		Mtermvectors:                       newMtermvectorsFunc(t),
		Ping:                               newPingFunc(t),
		ProfilingStacktraces:               newProfilingStacktracesFunc(t),
		ProfilingStatus:                    newProfilingStatusFunc(t),
		ProfilingTopnFunctions:             newProfilingTopnFunctionsFunc(t),
		PutScript:                          newPutScriptFunc(t),
		QueryRulesDeleteRule:               newQueryRulesDeleteRuleFunc(t),
		QueryRulesDeleteRuleset:            newQueryRulesDeleteRulesetFunc(t),
		QueryRulesGetRule:                  newQueryRulesGetRuleFunc(t),
		QueryRulesGetRuleset:               newQueryRulesGetRulesetFunc(t),
		QueryRulesListRulesets:             newQueryRulesListRulesetsFunc(t),
		QueryRulesPutRule:                  newQueryRulesPutRuleFunc(t),
		QueryRulesPutRuleset:               newQueryRulesPutRulesetFunc(t),
		QueryRulesTest:                     newQueryRulesTestFunc(t),
		RankEval:                           newRankEvalFunc(t),
		Reindex:                            newReindexFunc(t),
		ReindexRethrottle:                  newReindexRethrottleFunc(t),
		RenderSearchTemplate:               newRenderSearchTemplateFunc(t),
		ScriptsPainlessExecute:             newScriptsPainlessExecuteFunc(t),
		Scroll:                             newScrollFunc(t),
		SearchApplicationDeleteBehavioralAnalytics:    newSearchApplicationDeleteBehavioralAnalyticsFunc(t),
		SearchApplicationDelete:                       newSearchApplicationDeleteFunc(t),
		SearchApplicationGetBehavioralAnalytics:       newSearchApplicationGetBehavioralAnalyticsFunc(t),
		SearchApplicationGet:                          newSearchApplicationGetFunc(t),
		SearchApplicationList:                         newSearchApplicationListFunc(t),
		SearchApplicationPostBehavioralAnalyticsEvent: newSearchApplicationPostBehavioralAnalyticsEventFunc(t),
		SearchApplicationPutBehavioralAnalytics:       newSearchApplicationPutBehavioralAnalyticsFunc(t),
		SearchApplicationPut:                          newSearchApplicationPutFunc(t),
		SearchApplicationRenderQuery:                  newSearchApplicationRenderQueryFunc(t),
		SearchApplicationSearch:                       newSearchApplicationSearchFunc(t),
		SearchMvt:                                     newSearchMvtFunc(t),
		Search:                                        newSearchFunc(t),
		SearchShards:                                  newSearchShardsFunc(t),
		SearchTemplate:                                newSearchTemplateFunc(t),
		ShutdownDeleteNode:                            newShutdownDeleteNodeFunc(t),
		ShutdownGetNode:                               newShutdownGetNodeFunc(t),
		ShutdownPutNode:                               newShutdownPutNodeFunc(t),
		SimulateIngest:                                newSimulateIngestFunc(t),
		SynonymsDeleteSynonym:                         newSynonymsDeleteSynonymFunc(t),
		SynonymsDeleteSynonymRule:                     newSynonymsDeleteSynonymRuleFunc(t),
		SynonymsGetSynonym:                            newSynonymsGetSynonymFunc(t),
		SynonymsGetSynonymRule:                        newSynonymsGetSynonymRuleFunc(t),
		SynonymsGetSynonymsSets:                       newSynonymsGetSynonymsSetsFunc(t),
		SynonymsPutSynonym:                            newSynonymsPutSynonymFunc(t),
		SynonymsPutSynonymRule:                        newSynonymsPutSynonymRuleFunc(t),
		TermsEnum:                                     newTermsEnumFunc(t),
		Termvectors:                                   newTermvectorsFunc(t),
		UpdateByQuery:                                 newUpdateByQueryFunc(t),
		UpdateByQueryRethrottle:                       newUpdateByQueryRethrottleFunc(t),
		Update:                                        newUpdateFunc(t),
		Remote:                                        &Remote{},
	}
}
