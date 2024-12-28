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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827

package typedapi

import (
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	async_search_delete "github.com/elastic/go-elasticsearch/v8/typedapi/asyncsearch/delete"
	async_search_get "github.com/elastic/go-elasticsearch/v8/typedapi/asyncsearch/get"
	async_search_status "github.com/elastic/go-elasticsearch/v8/typedapi/asyncsearch/status"
	async_search_submit "github.com/elastic/go-elasticsearch/v8/typedapi/asyncsearch/submit"
	autoscaling_delete_autoscaling_policy "github.com/elastic/go-elasticsearch/v8/typedapi/autoscaling/deleteautoscalingpolicy"
	autoscaling_get_autoscaling_capacity "github.com/elastic/go-elasticsearch/v8/typedapi/autoscaling/getautoscalingcapacity"
	autoscaling_get_autoscaling_policy "github.com/elastic/go-elasticsearch/v8/typedapi/autoscaling/getautoscalingpolicy"
	autoscaling_put_autoscaling_policy "github.com/elastic/go-elasticsearch/v8/typedapi/autoscaling/putautoscalingpolicy"
	capabilities "github.com/elastic/go-elasticsearch/v8/typedapi/capabilities"
	ccr_delete_auto_follow_pattern "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/deleteautofollowpattern"
	ccr_follow "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/follow"
	ccr_follow_info "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/followinfo"
	ccr_follow_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/followstats"
	ccr_forget_follower "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/forgetfollower"
	ccr_get_auto_follow_pattern "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/getautofollowpattern"
	ccr_pause_auto_follow_pattern "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/pauseautofollowpattern"
	ccr_pause_follow "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/pausefollow"
	ccr_put_auto_follow_pattern "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/putautofollowpattern"
	ccr_resume_auto_follow_pattern "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/resumeautofollowpattern"
	ccr_resume_follow "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/resumefollow"
	ccr_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/stats"
	ccr_unfollow "github.com/elastic/go-elasticsearch/v8/typedapi/ccr/unfollow"
	connector_check_in "github.com/elastic/go-elasticsearch/v8/typedapi/connector/checkin"
	connector_delete "github.com/elastic/go-elasticsearch/v8/typedapi/connector/delete"
	connector_get "github.com/elastic/go-elasticsearch/v8/typedapi/connector/get"
	connector_last_sync "github.com/elastic/go-elasticsearch/v8/typedapi/connector/lastsync"
	connector_list "github.com/elastic/go-elasticsearch/v8/typedapi/connector/list"
	connector_post "github.com/elastic/go-elasticsearch/v8/typedapi/connector/post"
	connector_put "github.com/elastic/go-elasticsearch/v8/typedapi/connector/put"
	connector_secret_post "github.com/elastic/go-elasticsearch/v8/typedapi/connector/secretpost"
	connector_sync_job_cancel "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobcancel"
	connector_sync_job_delete "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobdelete"
	connector_sync_job_get "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobget"
	connector_sync_job_list "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjoblist"
	connector_sync_job_post "github.com/elastic/go-elasticsearch/v8/typedapi/connector/syncjobpost"
	connector_update_active_filtering "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateactivefiltering"
	connector_update_api_key_id "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateapikeyid"
	connector_update_configuration "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateconfiguration"
	connector_update_error "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateerror"
	connector_update_filtering "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatefiltering"
	connector_update_filtering_validation "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatefilteringvalidation"
	connector_update_index_name "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateindexname"
	connector_update_name "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatename"
	connector_update_native "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatenative"
	connector_update_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatepipeline"
	connector_update_scheduling "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatescheduling"
	connector_update_service_type "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updateservicetype"
	connector_update_status "github.com/elastic/go-elasticsearch/v8/typedapi/connector/updatestatus"
	core_bulk "github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	core_clear_scroll "github.com/elastic/go-elasticsearch/v8/typedapi/core/clearscroll"
	core_close_point_in_time "github.com/elastic/go-elasticsearch/v8/typedapi/core/closepointintime"
	core_count "github.com/elastic/go-elasticsearch/v8/typedapi/core/count"
	core_create "github.com/elastic/go-elasticsearch/v8/typedapi/core/create"
	core_delete "github.com/elastic/go-elasticsearch/v8/typedapi/core/delete"
	core_delete_by_query "github.com/elastic/go-elasticsearch/v8/typedapi/core/deletebyquery"
	core_delete_by_query_rethrottle "github.com/elastic/go-elasticsearch/v8/typedapi/core/deletebyqueryrethrottle"
	core_delete_script "github.com/elastic/go-elasticsearch/v8/typedapi/core/deletescript"
	core_exists "github.com/elastic/go-elasticsearch/v8/typedapi/core/exists"
	core_exists_source "github.com/elastic/go-elasticsearch/v8/typedapi/core/existssource"
	core_explain "github.com/elastic/go-elasticsearch/v8/typedapi/core/explain"
	core_field_caps "github.com/elastic/go-elasticsearch/v8/typedapi/core/fieldcaps"
	core_get "github.com/elastic/go-elasticsearch/v8/typedapi/core/get"
	core_get_script "github.com/elastic/go-elasticsearch/v8/typedapi/core/getscript"
	core_get_script_context "github.com/elastic/go-elasticsearch/v8/typedapi/core/getscriptcontext"
	core_get_script_languages "github.com/elastic/go-elasticsearch/v8/typedapi/core/getscriptlanguages"
	core_get_source "github.com/elastic/go-elasticsearch/v8/typedapi/core/getsource"
	core_health_report "github.com/elastic/go-elasticsearch/v8/typedapi/core/healthreport"
	core_index "github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
	core_info "github.com/elastic/go-elasticsearch/v8/typedapi/core/info"
	core_knn_search "github.com/elastic/go-elasticsearch/v8/typedapi/core/knnsearch"
	core_mget "github.com/elastic/go-elasticsearch/v8/typedapi/core/mget"
	core_msearch "github.com/elastic/go-elasticsearch/v8/typedapi/core/msearch"
	core_msearch_template "github.com/elastic/go-elasticsearch/v8/typedapi/core/msearchtemplate"
	core_mtermvectors "github.com/elastic/go-elasticsearch/v8/typedapi/core/mtermvectors"
	core_open_point_in_time "github.com/elastic/go-elasticsearch/v8/typedapi/core/openpointintime"
	core_ping "github.com/elastic/go-elasticsearch/v8/typedapi/core/ping"
	core_put_script "github.com/elastic/go-elasticsearch/v8/typedapi/core/putscript"
	core_rank_eval "github.com/elastic/go-elasticsearch/v8/typedapi/core/rankeval"
	core_reindex "github.com/elastic/go-elasticsearch/v8/typedapi/core/reindex"
	core_reindex_rethrottle "github.com/elastic/go-elasticsearch/v8/typedapi/core/reindexrethrottle"
	core_render_search_template "github.com/elastic/go-elasticsearch/v8/typedapi/core/rendersearchtemplate"
	core_scripts_painless_execute "github.com/elastic/go-elasticsearch/v8/typedapi/core/scriptspainlessexecute"
	core_scroll "github.com/elastic/go-elasticsearch/v8/typedapi/core/scroll"
	core_search "github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	core_search_mvt "github.com/elastic/go-elasticsearch/v8/typedapi/core/searchmvt"
	core_search_shards "github.com/elastic/go-elasticsearch/v8/typedapi/core/searchshards"
	core_search_template "github.com/elastic/go-elasticsearch/v8/typedapi/core/searchtemplate"
	core_terms_enum "github.com/elastic/go-elasticsearch/v8/typedapi/core/termsenum"
	core_termvectors "github.com/elastic/go-elasticsearch/v8/typedapi/core/termvectors"
	core_update "github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	core_update_by_query "github.com/elastic/go-elasticsearch/v8/typedapi/core/updatebyquery"
	core_update_by_query_rethrottle "github.com/elastic/go-elasticsearch/v8/typedapi/core/updatebyqueryrethrottle"
	dangling_indices_delete_dangling_index "github.com/elastic/go-elasticsearch/v8/typedapi/danglingindices/deletedanglingindex"
	dangling_indices_import_dangling_index "github.com/elastic/go-elasticsearch/v8/typedapi/danglingindices/importdanglingindex"
	dangling_indices_list_dangling_indices "github.com/elastic/go-elasticsearch/v8/typedapi/danglingindices/listdanglingindices"
	enrich_delete_policy "github.com/elastic/go-elasticsearch/v8/typedapi/enrich/deletepolicy"
	enrich_execute_policy "github.com/elastic/go-elasticsearch/v8/typedapi/enrich/executepolicy"
	enrich_get_policy "github.com/elastic/go-elasticsearch/v8/typedapi/enrich/getpolicy"
	enrich_put_policy "github.com/elastic/go-elasticsearch/v8/typedapi/enrich/putpolicy"
	enrich_stats "github.com/elastic/go-elasticsearch/v8/typedapi/enrich/stats"
	eql_delete "github.com/elastic/go-elasticsearch/v8/typedapi/eql/delete"
	eql_get "github.com/elastic/go-elasticsearch/v8/typedapi/eql/get"
	eql_get_status "github.com/elastic/go-elasticsearch/v8/typedapi/eql/getstatus"
	eql_search "github.com/elastic/go-elasticsearch/v8/typedapi/eql/search"
	esql_async_query "github.com/elastic/go-elasticsearch/v8/typedapi/esql/asyncquery"
	esql_query "github.com/elastic/go-elasticsearch/v8/typedapi/esql/query"
	features_get_features "github.com/elastic/go-elasticsearch/v8/typedapi/features/getfeatures"
	features_reset_features "github.com/elastic/go-elasticsearch/v8/typedapi/features/resetfeatures"
	fleet_global_checkpoints "github.com/elastic/go-elasticsearch/v8/typedapi/fleet/globalcheckpoints"
	fleet_msearch "github.com/elastic/go-elasticsearch/v8/typedapi/fleet/msearch"
	fleet_post_secret "github.com/elastic/go-elasticsearch/v8/typedapi/fleet/postsecret"
	fleet_search "github.com/elastic/go-elasticsearch/v8/typedapi/fleet/search"
	graph_explore "github.com/elastic/go-elasticsearch/v8/typedapi/graph/explore"
	ilm_delete_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/deletelifecycle"
	ilm_explain_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/explainlifecycle"
	ilm_get_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/getlifecycle"
	ilm_get_status "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/getstatus"
	ilm_migrate_to_data_tiers "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/migratetodatatiers"
	ilm_move_to_step "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/movetostep"
	ilm_put_lifecycle "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/putlifecycle"
	ilm_remove_policy "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/removepolicy"
	ilm_retry "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/retry"
	ilm_start "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/start"
	ilm_stop "github.com/elastic/go-elasticsearch/v8/typedapi/ilm/stop"
	inference_delete "github.com/elastic/go-elasticsearch/v8/typedapi/inference/delete"
	inference_get "github.com/elastic/go-elasticsearch/v8/typedapi/inference/get"
	inference_inference "github.com/elastic/go-elasticsearch/v8/typedapi/inference/inference"
	inference_put "github.com/elastic/go-elasticsearch/v8/typedapi/inference/put"
	ingest_delete_geoip_database "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/deletegeoipdatabase"
	ingest_delete_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/deletepipeline"
	ingest_geo_ip_stats "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/geoipstats"
	ingest_get_geoip_database "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/getgeoipdatabase"
	ingest_get_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/getpipeline"
	ingest_processor_grok "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/processorgrok"
	ingest_put_geoip_database "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/putgeoipdatabase"
	ingest_put_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/putpipeline"
	ingest_simulate "github.com/elastic/go-elasticsearch/v8/typedapi/ingest/simulate"
	license_delete "github.com/elastic/go-elasticsearch/v8/typedapi/license/delete"
	license_get "github.com/elastic/go-elasticsearch/v8/typedapi/license/get"
	license_get_basic_status "github.com/elastic/go-elasticsearch/v8/typedapi/license/getbasicstatus"
	license_get_trial_status "github.com/elastic/go-elasticsearch/v8/typedapi/license/gettrialstatus"
	license_post "github.com/elastic/go-elasticsearch/v8/typedapi/license/post"
	license_post_start_basic "github.com/elastic/go-elasticsearch/v8/typedapi/license/poststartbasic"
	license_post_start_trial "github.com/elastic/go-elasticsearch/v8/typedapi/license/poststarttrial"
	logstash_delete_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/logstash/deletepipeline"
	logstash_get_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/logstash/getpipeline"
	logstash_put_pipeline "github.com/elastic/go-elasticsearch/v8/typedapi/logstash/putpipeline"
	migration_deprecations "github.com/elastic/go-elasticsearch/v8/typedapi/migration/deprecations"
	migration_get_feature_upgrade_status "github.com/elastic/go-elasticsearch/v8/typedapi/migration/getfeatureupgradestatus"
	migration_post_feature_upgrade "github.com/elastic/go-elasticsearch/v8/typedapi/migration/postfeatureupgrade"
	profiling_flamegraph "github.com/elastic/go-elasticsearch/v8/typedapi/profiling/flamegraph"
	profiling_stacktraces "github.com/elastic/go-elasticsearch/v8/typedapi/profiling/stacktraces"
	profiling_status "github.com/elastic/go-elasticsearch/v8/typedapi/profiling/status"
	profiling_topn_functions "github.com/elastic/go-elasticsearch/v8/typedapi/profiling/topnfunctions"
	query_rules_delete_rule "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/deleterule"
	query_rules_delete_ruleset "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/deleteruleset"
	query_rules_get_rule "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/getrule"
	query_rules_get_ruleset "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/getruleset"
	query_rules_list_rulesets "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/listrulesets"
	query_rules_put_rule "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/putrule"
	query_rules_put_ruleset "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/putruleset"
	query_rules_test "github.com/elastic/go-elasticsearch/v8/typedapi/queryrules/test"
	rollup_delete_job "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/deletejob"
	rollup_get_jobs "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/getjobs"
	rollup_get_rollup_caps "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/getrollupcaps"
	rollup_put_job "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/putjob"
	rollup_rollup_search "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/rollupsearch"
	rollup_start_job "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/startjob"
	rollup_stop_job "github.com/elastic/go-elasticsearch/v8/typedapi/rollup/stopjob"
	search_application_delete "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/delete"
	search_application_delete_behavioral_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/deletebehavioralanalytics"
	search_application_get "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/get"
	search_application_get_behavioral_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/getbehavioralanalytics"
	search_application_list "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/list"
	search_application_put "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/put"
	search_application_put_behavioral_analytics "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/putbehavioralanalytics"
	search_application_search "github.com/elastic/go-elasticsearch/v8/typedapi/searchapplication/search"
	shutdown_delete_node "github.com/elastic/go-elasticsearch/v8/typedapi/shutdown/deletenode"
	shutdown_get_node "github.com/elastic/go-elasticsearch/v8/typedapi/shutdown/getnode"
	shutdown_put_node "github.com/elastic/go-elasticsearch/v8/typedapi/shutdown/putnode"
	sql_clear_cursor "github.com/elastic/go-elasticsearch/v8/typedapi/sql/clearcursor"
	sql_delete_async "github.com/elastic/go-elasticsearch/v8/typedapi/sql/deleteasync"
	sql_get_async "github.com/elastic/go-elasticsearch/v8/typedapi/sql/getasync"
	sql_get_async_status "github.com/elastic/go-elasticsearch/v8/typedapi/sql/getasyncstatus"
	sql_query "github.com/elastic/go-elasticsearch/v8/typedapi/sql/query"
	sql_translate "github.com/elastic/go-elasticsearch/v8/typedapi/sql/translate"
	ssl_certificates "github.com/elastic/go-elasticsearch/v8/typedapi/ssl/certificates"
	synonyms_delete_synonym "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/deletesynonym"
	synonyms_delete_synonym_rule "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/deletesynonymrule"
	synonyms_get_synonym "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/getsynonym"
	synonyms_get_synonym_rule "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/getsynonymrule"
	synonyms_get_synonyms_sets "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/getsynonymssets"
	synonyms_put_synonym "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/putsynonym"
	synonyms_put_synonym_rule "github.com/elastic/go-elasticsearch/v8/typedapi/synonyms/putsynonymrule"
	text_structure_find_field_structure "github.com/elastic/go-elasticsearch/v8/typedapi/textstructure/findfieldstructure"
	text_structure_find_message_structure "github.com/elastic/go-elasticsearch/v8/typedapi/textstructure/findmessagestructure"
	text_structure_find_structure "github.com/elastic/go-elasticsearch/v8/typedapi/textstructure/findstructure"
	text_structure_test_grok_pattern "github.com/elastic/go-elasticsearch/v8/typedapi/textstructure/testgrokpattern"
	transform_delete_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/deletetransform"
	transform_get_node_stats "github.com/elastic/go-elasticsearch/v8/typedapi/transform/getnodestats"
	transform_get_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/gettransform"
	transform_get_transform_stats "github.com/elastic/go-elasticsearch/v8/typedapi/transform/gettransformstats"
	transform_preview_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/previewtransform"
	transform_put_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/puttransform"
	transform_reset_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/resettransform"
	transform_schedule_now_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/schedulenowtransform"
	transform_start_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/starttransform"
	transform_stop_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/stoptransform"
	transform_update_transform "github.com/elastic/go-elasticsearch/v8/typedapi/transform/updatetransform"
	transform_upgrade_transforms "github.com/elastic/go-elasticsearch/v8/typedapi/transform/upgradetransforms"
	watcher_ack_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/ackwatch"
	watcher_activate_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/activatewatch"
	watcher_deactivate_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/deactivatewatch"
	watcher_delete_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/deletewatch"
	watcher_execute_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/executewatch"
	watcher_get_settings "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/getsettings"
	watcher_get_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/getwatch"
	watcher_put_watch "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/putwatch"
	watcher_query_watches "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/querywatches"
	watcher_start "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/start"
	watcher_stats "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/stats"
	watcher_stop "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/stop"
	watcher_update_settings "github.com/elastic/go-elasticsearch/v8/typedapi/watcher/updatesettings"
)

type AsyncSearch struct {
	// Delete an async search.
	//
	// If the asynchronous search is still running, it is cancelled.
	// Otherwise, the saved search results are deleted.
	// If the Elasticsearch security features are enabled, the deletion of a
	// specific async search is restricted to: the authenticated user that submitted
	// the original search request; users that have the `cancel_task` cluster
	// privilege.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Delete async_search_delete.NewDelete
	// Get async search results.
	//
	// Retrieve the results of a previously submitted asynchronous search request.
	// If the Elasticsearch security features are enabled, access to the results of
	// a specific async search is restricted to the user or API key that submitted
	// it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Get async_search_get.NewGet
	// Get the async search status.
	//
	// Get the status of a previously submitted async search request given its
	// identifier, without retrieving search results.
	// If the Elasticsearch security features are enabled, use of this API is
	// restricted to the `monitoring_user` role.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Status async_search_status.NewStatus
	// Run an async search.
	//
	// When the primary sort of the results is an indexed field, shards get sorted
	// based on minimum and maximum value that they hold for that field. Partial
	// results become available following the sort criteria that was requested.
	//
	// Warning: Asynchronous search does not support scroll or search requests that
	// include only the suggest section.
	//
	// By default, Elasticsearch does not allow you to store an async search
	// response larger than 10Mb and an attempt to do this results in an error.
	// The maximum allowed size for a stored async search response can be set by
	// changing the `search.max_async_search_response_size` cluster level setting.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html
	Submit async_search_submit.NewSubmit
}

type Autoscaling struct {
	// Delete an autoscaling policy.
	//
	// NOTE: This feature is designed for indirect use by Elasticsearch Service,
	// Elastic Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
	// supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-delete-autoscaling-policy.html
	DeleteAutoscalingPolicy autoscaling_delete_autoscaling_policy.NewDeleteAutoscalingPolicy
	// Get the autoscaling capacity.
	//
	// NOTE: This feature is designed for indirect use by Elasticsearch Service,
	// Elastic Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
	// supported.
	//
	// This API gets the current autoscaling capacity based on the configured
	// autoscaling policy.
	// It will return information to size the cluster appropriately to the current
	// workload.
	//
	// The `required_capacity` is calculated as the maximum of the
	// `required_capacity` result of all individual deciders that are enabled for
	// the policy.
	//
	// The operator should verify that the `current_nodes` match the operator’s
	// knowledge of the cluster to avoid making autoscaling decisions based on stale
	// or incomplete information.
	//
	// The response contains decider-specific information you can use to diagnose
	// how and why autoscaling determined a certain capacity was required.
	// This information is provided for diagnosis only.
	// Do not use this information to make autoscaling decisions.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-get-autoscaling-capacity.html
	GetAutoscalingCapacity autoscaling_get_autoscaling_capacity.NewGetAutoscalingCapacity
	// Get an autoscaling policy.
	//
	// NOTE: This feature is designed for indirect use by Elasticsearch Service,
	// Elastic Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
	// supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-get-autoscaling-capacity.html
	GetAutoscalingPolicy autoscaling_get_autoscaling_policy.NewGetAutoscalingPolicy
	// Create or update an autoscaling policy.
	//
	// NOTE: This feature is designed for indirect use by Elasticsearch Service,
	// Elastic Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
	// supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-put-autoscaling-policy.html
	PutAutoscalingPolicy autoscaling_put_autoscaling_policy.NewPutAutoscalingPolicy
}

type Capabilities struct {
	// Checks if the specified combination of method, API, parameters, and arbitrary
	// capabilities are supported
	// https://github.com/elastic/elasticsearch/blob/main/rest-api-spec/src/yamlRestTest/resources/rest-api-spec/test/README.asciidoc#require-or-skip-api-capabilities
	Capabilities capabilities.NewCapabilities
}

type Ccr struct {
	// Deletes auto-follow patterns.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-delete-auto-follow-pattern.html
	DeleteAutoFollowPattern ccr_delete_auto_follow_pattern.NewDeleteAutoFollowPattern
	// Creates a new follower index configured to follow the referenced leader
	// index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-put-follow.html
	Follow ccr_follow.NewFollow
	// Retrieves information about all follower indices, including parameters and
	// status for each follower index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-follow-info.html
	FollowInfo ccr_follow_info.NewFollowInfo
	// Retrieves follower stats. return shard-level stats about the following tasks
	// associated with each shard for the specified indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-follow-stats.html
	FollowStats ccr_follow_stats.NewFollowStats
	// Removes the follower retention leases from the leader.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-forget-follower.html
	ForgetFollower ccr_forget_follower.NewForgetFollower
	// Gets configured auto-follow patterns. Returns the specified auto-follow
	// pattern collection.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-auto-follow-pattern.html
	GetAutoFollowPattern ccr_get_auto_follow_pattern.NewGetAutoFollowPattern
	// Pauses an auto-follow pattern
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-pause-auto-follow-pattern.html
	PauseAutoFollowPattern ccr_pause_auto_follow_pattern.NewPauseAutoFollowPattern
	// Pauses a follower index. The follower index will not fetch any additional
	// operations from the leader index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-pause-follow.html
	PauseFollow ccr_pause_follow.NewPauseFollow
	// Creates a new named collection of auto-follow patterns against a specified
	// remote cluster. Newly created indices on the remote cluster matching any of
	// the specified patterns will be automatically configured as follower indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-put-auto-follow-pattern.html
	PutAutoFollowPattern ccr_put_auto_follow_pattern.NewPutAutoFollowPattern
	// Resumes an auto-follow pattern that has been paused
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-resume-auto-follow-pattern.html
	ResumeAutoFollowPattern ccr_resume_auto_follow_pattern.NewResumeAutoFollowPattern
	// Resumes a follower index that has been paused
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-resume-follow.html
	ResumeFollow ccr_resume_follow.NewResumeFollow
	// Gets all stats related to cross-cluster replication.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-get-stats.html
	Stats ccr_stats.NewStats
	// Stops the following task associated with a follower index and removes index
	// metadata and settings associated with cross-cluster replication.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-post-unfollow.html
	Unfollow ccr_unfollow.NewUnfollow
}

type Connector struct {
	// Check in a connector.
	//
	// Update the `last_seen` field in the connector and set it to the current
	// timestamp.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/check-in-connector-api.html
	CheckIn connector_check_in.NewCheckIn
	// Delete a connector.
	//
	// Removes a connector and associated sync jobs.
	// This is a destructive action that is not recoverable.
	// NOTE: This action doesn’t delete any API keys, ingest pipelines, or data
	// indices associated with the connector.
	// These need to be removed manually.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-connector-api.html
	Delete connector_delete.NewDelete
	// Get a connector.
	//
	// Get the details about a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-connector-api.html
	Get connector_get.NewGet
	// Update the connector last sync stats.
	//
	// Update the fields related to the last sync of a connector.
	// This action is used for analytics and monitoring.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-last-sync-api.html
	LastSync connector_last_sync.NewLastSync
	// Get all connectors.
	//
	// Get information about all connectors.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-connector-api.html
	List connector_list.NewList
	// Create a connector.
	//
	// Connectors are Elasticsearch integrations that bring content from third-party
	// data sources, which can be deployed on Elastic Cloud or hosted on your own
	// infrastructure.
	// Elastic managed connectors (Native connectors) are a managed service on
	// Elastic Cloud.
	// Self-managed connectors (Connector clients) are self-managed on your
	// infrastructure.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/create-connector-api.html
	Post connector_post.NewPost
	// Create or update a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/create-connector-api.html
	Put connector_put.NewPut
	// Creates a secret for a Connector.
	//
	SecretPost connector_secret_post.NewSecretPost
	// Cancel a connector sync job.
	//
	// Cancel a connector sync job, which sets the status to cancelling and updates
	// `cancellation_requested_at` to the current time.
	// The connector service is then responsible for setting the status of connector
	// sync jobs to cancelled.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/cancel-connector-sync-job-api.html
	SyncJobCancel connector_sync_job_cancel.NewSyncJobCancel
	// Delete a connector sync job.
	//
	// Remove a connector sync job and its associated data.
	// This is a destructive action that is not recoverable.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-connector-sync-job-api.html
	SyncJobDelete connector_sync_job_delete.NewSyncJobDelete
	// Get a connector sync job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-connector-sync-job-api.html
	SyncJobGet connector_sync_job_get.NewSyncJobGet
	// Get all connector sync jobs.
	//
	// Get information about all stored connector sync jobs listed by their creation
	// date in ascending order.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-connector-sync-jobs-api.html
	SyncJobList connector_sync_job_list.NewSyncJobList
	// Create a connector sync job.
	//
	// Create a connector sync job document in the internal index and initialize its
	// counters and timestamps with default values.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/create-connector-sync-job-api.html
	SyncJobPost connector_sync_job_post.NewSyncJobPost
	// Activate the connector draft filter.
	//
	// Activates the valid draft filtering for a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-filtering-api.html
	UpdateActiveFiltering connector_update_active_filtering.NewUpdateActiveFiltering
	// Update the connector API key ID.
	//
	// Update the `api_key_id` and `api_key_secret_id` fields of a connector.
	// You can specify the ID of the API key used for authorization and the ID of
	// the connector secret where the API key is stored.
	// The connector secret ID is required only for Elastic managed (native)
	// connectors.
	// Self-managed connectors (connector clients) do not use this field.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-api-key-id-api.html
	UpdateApiKeyId connector_update_api_key_id.NewUpdateApiKeyId
	// Update the connector configuration.
	//
	// Update the configuration field in the connector document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-configuration-api.html
	UpdateConfiguration connector_update_configuration.NewUpdateConfiguration
	// Update the connector error field.
	//
	// Set the error field for the connector.
	// If the error provided in the request body is non-null, the connector’s status
	// is updated to error.
	// Otherwise, if the error is reset to null, the connector status is updated to
	// connected.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-error-api.html
	UpdateError connector_update_error.NewUpdateError
	// Update the connector filtering.
	//
	// Update the draft filtering configuration of a connector and marks the draft
	// validation state as edited.
	// The filtering draft is activated once validated by the running Elastic
	// connector service.
	// The filtering property is used to configure sync rules (both basic and
	// advanced) for a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-filtering-api.html
	UpdateFiltering connector_update_filtering.NewUpdateFiltering
	// Update the connector draft filtering validation.
	//
	// Update the draft filtering validation info for a connector.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-filtering-validation-api.html
	UpdateFilteringValidation connector_update_filtering_validation.NewUpdateFilteringValidation
	// Update the connector index name.
	//
	// Update the `index_name` field of a connector, specifying the index where the
	// data ingested by the connector is stored.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-index-name-api.html
	UpdateIndexName connector_update_index_name.NewUpdateIndexName
	// Update the connector name and description.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-name-description-api.html
	UpdateName connector_update_name.NewUpdateName
	// Update the connector is_native flag.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-native-api.html
	UpdateNative connector_update_native.NewUpdateNative
	// Update the connector pipeline.
	//
	// When you create a new connector, the configuration of an ingest pipeline is
	// populated with default settings.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-pipeline-api.html
	UpdatePipeline connector_update_pipeline.NewUpdatePipeline
	// Update the connector scheduling.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-scheduling-api.html
	UpdateScheduling connector_update_scheduling.NewUpdateScheduling
	// Update the connector service type.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-service-type-api.html
	UpdateServiceType connector_update_service_type.NewUpdateServiceType
	// Update the connector status.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-connector-status-api.html
	UpdateStatus connector_update_status.NewUpdateStatus
}

type Core struct {
	// Bulk index or delete documents.
	// Performs multiple indexing or delete operations in a single API call.
	// This reduces overhead and can greatly increase indexing speed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
	Bulk core_bulk.NewBulk
	// Clear a scrolling search.
	//
	// Clear the search context and results for a scrolling search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-scroll-api.html
	ClearScroll core_clear_scroll.NewClearScroll
	// Close a point in time.
	//
	// A point in time must be opened explicitly before being used in search
	// requests.
	// The `keep_alive` parameter tells Elasticsearch how long it should persist.
	// A point in time is automatically closed when the `keep_alive` period has
	// elapsed.
	// However, keeping points in time has a cost; close them as soon as they are no
	// longer required for search requests.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	ClosePointInTime core_close_point_in_time.NewClosePointInTime
	// Returns number of documents matching a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
	Count core_count.NewCount
	// Index a document.
	// Adds a JSON document to the specified data stream or index and makes it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Create core_create.NewCreate
	// Delete a document.
	// Removes a JSON document from the specified index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
	Delete core_delete.NewDelete
	// Delete documents.
	// Deletes documents that match the specified query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQuery core_delete_by_query.NewDeleteByQuery
	// Throttle a delete by query operation.
	//
	// Change the number of requests per second for a particular delete by query
	// operation.
	// Rethrottling that speeds up the query takes effect immediately but
	// rethrotting that slows down the query takes effect after completing the
	// current batch to prevent scroll timeouts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQueryRethrottle core_delete_by_query_rethrottle.NewDeleteByQueryRethrottle
	// Delete a script or search template.
	// Deletes a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	DeleteScript core_delete_script.NewDeleteScript
	// Check a document.
	// Checks if a specified document exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Exists core_exists.NewExists
	// Check for a document source.
	// Checks if a document's `_source` is stored.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	ExistsSource core_exists_source.NewExistsSource
	// Explain a document match result.
	// Returns information about why a specific document matches, or doesn’t match,
	// a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
	Explain core_explain.NewExplain
	// Get the field capabilities.
	//
	// Get information about the capabilities of fields among multiple indices.
	//
	// For data streams, the API returns field capabilities among the stream’s
	// backing indices.
	// It returns runtime fields like any other field.
	// For example, a runtime field with a type of keyword is returned the same as
	// any other field that belongs to the `keyword` family.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-field-caps.html
	FieldCaps core_field_caps.NewFieldCaps
	// Get a document by its ID.
	// Retrieves the document with the specified ID from an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Get core_get.NewGet
	// Get a script or search template.
	// Retrieves a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	GetScript core_get_script.NewGetScript
	// Get script contexts.
	//
	// Get a list of supported script contexts and their methods.
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-contexts.html
	GetScriptContext core_get_script_context.NewGetScriptContext
	// Get script languages.
	//
	// Get a list of available script types, languages, and contexts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	GetScriptLanguages core_get_script_languages.NewGetScriptLanguages
	// Get a document's source.
	// Returns the source of a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	GetSource core_get_source.NewGetSource
	// Returns the health of the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/health-api.html
	HealthReport core_health_report.NewHealthReport
	// Index a document.
	// Adds a JSON document to the specified data stream or index and makes it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Index core_index.NewIndex
	// Get cluster info.
	// Returns basic information about the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Info core_info.NewInfo
	// Run a knn search.
	//
	// NOTE: The kNN search API has been replaced by the `knn` option in the search
	// API.
	//
	// Perform a k-nearest neighbor (kNN) search on a dense_vector field and return
	// the matching documents.
	// Given a query vector, the API finds the k closest vectors and returns those
	// documents as search hits.
	//
	// Elasticsearch uses the HNSW algorithm to support efficient kNN search.
	// Like most kNN algorithms, HNSW is an approximate method that sacrifices
	// result accuracy for improved search speed.
	// This means the results returned are not always the true k closest neighbors.
	//
	// The kNN search API supports restricting the search using a filter.
	// The search will return the top k documents that also match the filter query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	KnnSearch core_knn_search.NewKnnSearch
	// Get multiple documents.
	//
	// Get multiple JSON documents by ID from one or more indices.
	// If you specify an index in the request URI, you only need to specify the
	// document IDs in the request body.
	// To ensure fast responses, this multi get (mget) API responds with partial
	// results if one or more shards fail.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-get.html
	Mget core_mget.NewMget
	// Run multiple searches.
	//
	// The format of the request is similar to the bulk API format and makes use of
	// the newline delimited JSON (NDJSON) format.
	// The structure is as follows:
	//
	// ```
	// header\n
	// body\n
	// header\n
	// body\n
	// ```
	//
	// This structure is specifically optimized to reduce parsing if a specific
	// search ends up redirected to another node.
	//
	// IMPORTANT: The final line of data must end with a newline character `\n`.
	// Each newline character may be preceded by a carriage return `\r`.
	// When sending requests to this endpoint the `Content-Type` header should be
	// set to `application/x-ndjson`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	Msearch core_msearch.NewMsearch
	// Run multiple templated searches.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	MsearchTemplate core_msearch_template.NewMsearchTemplate
	// Get multiple term vectors.
	//
	// You can specify existing documents by index and ID or provide artificial
	// documents in the body of the request.
	// You can specify the index in the request body or request URI.
	// The response contains a `docs` array with all the fetched termvectors.
	// Each element has the structure provided by the termvectors API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
	Mtermvectors core_mtermvectors.NewMtermvectors
	// Open a point in time.
	//
	// A search request by default runs against the most recent visible data of the
	// target indices,
	// which is called point in time. Elasticsearch pit (point in time) is a
	// lightweight view into the
	// state of the data as it existed when initiated. In some cases, it’s preferred
	// to perform multiple
	// search requests using the same point in time. For example, if refreshes
	// happen between
	// `search_after` requests, then the results of those requests might not be
	// consistent as changes happening
	// between searches are only visible to the more recent point in time.
	//
	// A point in time must be opened explicitly before being used in search
	// requests.
	// The `keep_alive` parameter tells Elasticsearch how long it should persist.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	OpenPointInTime core_open_point_in_time.NewOpenPointInTime
	// Ping the cluster.
	// Returns whether the cluster is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Ping core_ping.NewPing
	// Create or update a script or search template.
	// Creates or updates a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	PutScript core_put_script.NewPutScript
	// Evaluate ranked search results.
	//
	// Evaluate the quality of ranked search results over a set of typical search
	// queries.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
	RankEval core_rank_eval.NewRankEval
	// Reindex documents.
	// Copies documents from a source to a destination. The source can be any
	// existing index, alias, or data stream. The destination must differ from the
	// source. For example, you cannot reindex a data stream into itself.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	Reindex core_reindex.NewReindex
	// Throttle a reindex operation.
	//
	// Change the number of requests per second for a particular reindex operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	ReindexRethrottle core_reindex_rethrottle.NewReindexRethrottle
	// Render a search template.
	//
	// Render a search template as a search request body.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/render-search-template-api.html
	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate
	// Run a script.
	// Runs a script and returns a result.
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute
	// Run a scrolling search.
	//
	// IMPORTANT: The scroll API is no longer recommend for deep pagination. If you
	// need to preserve the index state while paging through more than 10,000 hits,
	// use the `search_after` parameter with a point in time (PIT).
	//
	// The scroll API gets large sets of results from a single scrolling search
	// request.
	// To get the necessary scroll ID, submit a search API request that includes an
	// argument for the `scroll` query parameter.
	// The `scroll` parameter indicates how long Elasticsearch should retain the
	// search context for the request.
	// The search response returns a scroll ID in the `_scroll_id` response body
	// parameter.
	// You can then use the scroll ID with the scroll API to retrieve the next batch
	// of results for the request.
	// If the Elasticsearch security features are enabled, the access to the results
	// of a specific scroll ID is restricted to the user or API key that submitted
	// the search.
	//
	// You can also use the scroll API to specify a new scroll parameter that
	// extends or shortens the retention period for the search context.
	//
	// IMPORTANT: Results from a scrolling search reflect the state of the index at
	// the time of the initial search request. Subsequent indexing or document
	// changes only affect later search and scroll requests.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-body.html#request-body-search-scroll
	Scroll core_scroll.NewScroll
	// Run a search.
	//
	// Get search hits that match the query defined in the request.
	// You can provide search queries using the `q` query string parameter or the
	// request body.
	// If both are specified, only the query parameter is used.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	Search core_search.NewSearch
	// Search a vector tile.
	//
	// Search a vector tile for geospatial values.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-vector-tile-api.html
	SearchMvt core_search_mvt.NewSearchMvt
	// Get the search shards.
	//
	// Get the indices and shards that a search request would be run against.
	// This information can be useful for working out issues or planning
	// optimizations with routing and shard preferences.
	// When filtered aliases are used, the filter is returned as part of the indices
	// section.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-shards.html
	SearchShards core_search_shards.NewSearchShards
	// Run a search with a search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template.html
	SearchTemplate core_search_template.NewSearchTemplate
	// Get terms in an index.
	//
	// Discover terms that match a partial string in an index.
	// This "terms enum" API is designed for low-latency look-ups used in
	// auto-complete scenarios.
	//
	// If the `complete` property in the response is false, the returned terms set
	// may be incomplete and should be treated as approximate.
	// This can occur due to a few reasons, such as a request timeout or a node
	// error.
	//
	// NOTE: The terms enum API may return terms from deleted documents. Deleted
	// documents are initially only marked as deleted. It is not until their
	// segments are merged that documents are actually deleted. Until that happens,
	// the terms enum API will return terms from these documents.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-terms-enum.html
	TermsEnum core_terms_enum.NewTermsEnum
	// Get term vector information.
	//
	// Get information and statistics about terms in the fields of a particular
	// document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-termvectors.html
	Termvectors core_termvectors.NewTermvectors
	// Update a document.
	// Updates a document by running a script or passing a partial document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update.html
	Update core_update.NewUpdate
	// Update documents.
	// Updates documents that match the specified query.
	// If no query is specified, performs an update on every document in the data
	// stream or index without modifying the source, which is useful for picking up
	// mapping changes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
	UpdateByQuery core_update_by_query.NewUpdateByQuery
	// Throttle an update by query operation.
	//
	// Change the number of requests per second for a particular update by query
	// operation.
	// Rethrottling that speeds up the query takes effect immediately but
	// rethrotting that slows down the query takes effect after completing the
	// current batch to prevent scroll timeouts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
	UpdateByQueryRethrottle core_update_by_query_rethrottle.NewUpdateByQueryRethrottle
}

type DanglingIndices struct {
	// Delete a dangling index.
	//
	// If Elasticsearch encounters index data that is absent from the current
	// cluster state, those indices are considered to be dangling.
	// For example, this can happen if you delete more than
	// `cluster.indices.tombstones.size` indices while an Elasticsearch node is
	// offline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-gateway-dangling-indices.html
	DeleteDanglingIndex dangling_indices_delete_dangling_index.NewDeleteDanglingIndex
	// Import a dangling index.
	//
	// If Elasticsearch encounters index data that is absent from the current
	// cluster state, those indices are considered to be dangling.
	// For example, this can happen if you delete more than
	// `cluster.indices.tombstones.size` indices while an Elasticsearch node is
	// offline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-gateway-dangling-indices.html
	ImportDanglingIndex dangling_indices_import_dangling_index.NewImportDanglingIndex
	// Get the dangling indices.
	//
	// If Elasticsearch encounters index data that is absent from the current
	// cluster state, those indices are considered to be dangling.
	// For example, this can happen if you delete more than
	// `cluster.indices.tombstones.size` indices while an Elasticsearch node is
	// offline.
	//
	// Use this API to list dangling indices, which you can then import or delete.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-gateway-dangling-indices.html
	ListDanglingIndices dangling_indices_list_dangling_indices.NewListDanglingIndices
}

type Enrich struct {
	// Delete an enrich policy.
	// Deletes an existing enrich policy and its enrich index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-enrich-policy-api.html
	DeletePolicy enrich_delete_policy.NewDeletePolicy
	// Creates the enrich index for an existing enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/execute-enrich-policy-api.html
	ExecutePolicy enrich_execute_policy.NewExecutePolicy
	// Get an enrich policy.
	// Returns information about an enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-enrich-policy-api.html
	GetPolicy enrich_get_policy.NewGetPolicy
	// Create an enrich policy.
	// Creates an enrich policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-enrich-policy-api.html
	PutPolicy enrich_put_policy.NewPutPolicy
	// Get enrich stats.
	// Returns enrich coordinator statistics and information about enrich policies
	// that are currently executing.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/enrich-stats-api.html
	Stats enrich_stats.NewStats
}

type Eql struct {
	// Deletes an async EQL search or a stored synchronous EQL search.
	// The API also deletes results for the search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
	Delete eql_delete.NewDelete
	// Returns the current status and available results for an async EQL search or a
	// stored synchronous EQL search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-eql-search-api.html
	Get eql_get.NewGet
	// Returns the current status for an async EQL search or a stored synchronous
	// EQL search without returning results.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-eql-status-api.html
	GetStatus eql_get_status.NewGetStatus
	// Returns results matching a query expressed in Event Query Language (EQL)
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
	Search eql_search.NewSearch
}

type Esql struct {
	// Executes an ESQL request asynchronously
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-async-query-api.html
	AsyncQuery esql_async_query.NewAsyncQuery
	// Executes an ES|QL request
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-rest.html
	Query esql_query.NewQuery
}

type Features struct {
	// Gets a list of features which can be included in snapshots using the
	// feature_states field when creating a snapshot
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-features-api.html
	GetFeatures features_get_features.NewGetFeatures
	// Resets the internal state of features, usually by deleting system indices
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html
	ResetFeatures features_reset_features.NewResetFeatures
}

type Fleet struct {
	// Returns the current global checkpoints for an index. This API is design for
	// internal use by the fleet server project.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-global-checkpoints.html
	GlobalCheckpoints fleet_global_checkpoints.NewGlobalCheckpoints
	// Executes several [fleet
	// searches](https://www.elastic.co/guide/en/elasticsearch/reference/current/fleet-search.html)
	// with a single API request.
	// The API follows the same structure as the [multi
	// search](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html)
	// API. However, similar to the fleet search API, it
	// supports the wait_for_checkpoints parameter.
	//
	Msearch fleet_msearch.NewMsearch
	// Creates a secret stored by Fleet.
	//
	PostSecret fleet_post_secret.NewPostSecret
	// The purpose of the fleet search api is to provide a search api where the
	// search will only be executed
	// after provided checkpoint has been processed and is visible for searches
	// inside of Elasticsearch.
	//
	Search fleet_search.NewSearch
}

type Graph struct {
	// Extracts and summarizes information about the documents and terms in an
	// Elasticsearch data stream or index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/graph-explore-api.html
	Explore graph_explore.NewExplore
}

type Ilm struct {
	// Deletes the specified lifecycle policy definition. You cannot delete policies
	// that are currently in use. If the policy is being used to manage any indices,
	// the request fails and returns an error.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-delete-lifecycle.html
	DeleteLifecycle ilm_delete_lifecycle.NewDeleteLifecycle
	// Retrieves information about the index’s current lifecycle state, such as the
	// currently executing phase, action, and step. Shows when the index entered
	// each one, the definition of the running phase, and information about any
	// failures.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-explain-lifecycle.html
	ExplainLifecycle ilm_explain_lifecycle.NewExplainLifecycle
	// Retrieves a lifecycle policy.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-get-lifecycle.html
	GetLifecycle ilm_get_lifecycle.NewGetLifecycle
	// Retrieves the current index lifecycle management (ILM) status.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-get-status.html
	GetStatus ilm_get_status.NewGetStatus
	// Switches the indices, ILM policies, and legacy, composable and component
	// templates from using custom node attributes and
	// attribute-based allocation filters to using data tiers, and optionally
	// deletes one legacy index template.+
	// Using node roles enables ILM to automatically move the indices between data
	// tiers.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-migrate-to-data-tiers.html
	MigrateToDataTiers ilm_migrate_to_data_tiers.NewMigrateToDataTiers
	// Manually moves an index into the specified step and executes that step.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-move-to-step.html
	MoveToStep ilm_move_to_step.NewMoveToStep
	// Creates a lifecycle policy. If the specified policy exists, the policy is
	// replaced and the policy version is incremented.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-put-lifecycle.html
	PutLifecycle ilm_put_lifecycle.NewPutLifecycle
	// Removes the assigned lifecycle policy and stops managing the specified index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-remove-policy.html
	RemovePolicy ilm_remove_policy.NewRemovePolicy
	// Retries executing the policy for an index that is in the ERROR step.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-retry-policy.html
	Retry ilm_retry.NewRetry
	// Start the index lifecycle management (ILM) plugin.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-start.html
	Start ilm_start.NewStart
	// Halts all lifecycle management operations and stops the index lifecycle
	// management (ILM) plugin
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-stop.html
	Stop ilm_stop.NewStop
}

type Inference struct {
	// Delete an inference endpoint
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-inference-api.html
	Delete inference_delete.NewDelete
	// Get an inference endpoint
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-inference-api.html
	Get inference_get.NewGet
	// Perform inference on the service
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/post-inference-api.html
	Inference inference_inference.NewInference
	// Create an inference endpoint
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-inference-api.html
	Put inference_put.NewPut
}

type Ingest struct {
	// Deletes a geoip database configuration.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-geoip-database-api.html
	DeleteGeoipDatabase ingest_delete_geoip_database.NewDeleteGeoipDatabase
	// Deletes one or more existing ingest pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-pipeline-api.html
	DeletePipeline ingest_delete_pipeline.NewDeletePipeline
	// Gets download statistics for GeoIP2 databases used with the geoip processor.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html
	GeoIpStats ingest_geo_ip_stats.NewGeoIpStats
	// Returns information about one or more geoip database configurations.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-geoip-database-api.html
	GetGeoipDatabase ingest_get_geoip_database.NewGetGeoipDatabase
	// Returns information about one or more ingest pipelines.
	// This API returns a local reference of the pipeline.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-pipeline-api.html
	GetPipeline ingest_get_pipeline.NewGetPipeline
	// Extracts structured fields out of a single text field within a document.
	// You choose which field to extract matched fields from, as well as the grok
	// pattern you expect will match.
	// A grok pattern is like a regular expression that supports aliased expressions
	// that can be reused.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/grok-processor.html
	ProcessorGrok ingest_processor_grok.NewProcessorGrok
	// Returns information about one or more geoip database configurations.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-geoip-database-api.html
	PutGeoipDatabase ingest_put_geoip_database.NewPutGeoipDatabase
	// Creates or updates an ingest pipeline.
	// Changes made using this API take effect immediately.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html
	PutPipeline ingest_put_pipeline.NewPutPipeline
	// Executes an ingest pipeline against a set of provided documents.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/simulate-pipeline-api.html
	Simulate ingest_simulate.NewSimulate
}

type License struct {
	// Deletes licensing information for the cluster
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-license.html
	Delete license_delete.NewDelete
	// Get license information.
	// Returns information about your Elastic license, including its type, its
	// status, when it was issued, and when it expires.
	// For more information about the different types of licenses, refer to [Elastic
	// Stack subscriptions](https://www.elastic.co/subscriptions).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-license.html
	Get license_get.NewGet
	// Retrieves information about the status of the basic license.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-basic-status.html
	GetBasicStatus license_get_basic_status.NewGetBasicStatus
	// Retrieves information about the status of the trial license.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-trial-status.html
	GetTrialStatus license_get_trial_status.NewGetTrialStatus
	// Updates the license for the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-license.html
	Post license_post.NewPost
	// The start basic API enables you to initiate an indefinite basic license,
	// which gives access to all the basic features. If the basic license does not
	// support all of the features that are available with your current license,
	// however, you are notified in the response. You must then re-submit the API
	// request with the acknowledge parameter set to true.
	// To check the status of your basic license, use the following API: [Get basic
	// status](https://www.elastic.co/guide/en/elasticsearch/reference/current/get-basic-status.html).
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-basic.html
	PostStartBasic license_post_start_basic.NewPostStartBasic
	// The start trial API enables you to start a 30-day trial, which gives access
	// to all subscription features.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-trial.html
	PostStartTrial license_post_start_trial.NewPostStartTrial
}

type Logstash struct {
	// Deletes a pipeline used for Logstash Central Management.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-delete-pipeline.html
	DeletePipeline logstash_delete_pipeline.NewDeletePipeline
	// Retrieves pipelines used for Logstash Central Management.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-get-pipeline.html
	GetPipeline logstash_get_pipeline.NewGetPipeline
	// Creates or updates a pipeline used for Logstash Central Management.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-api-put-pipeline.html
	PutPipeline logstash_put_pipeline.NewPutPipeline
}

type Migration struct {
	// Retrieves information about different cluster, node, and index level settings
	// that use deprecated features that will be removed or changed in the next
	// major version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/migration-api-deprecation.html
	Deprecations migration_deprecations.NewDeprecations
	// Find out whether system features need to be upgraded or not
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/migration-api-feature-upgrade.html
	GetFeatureUpgradeStatus migration_get_feature_upgrade_status.NewGetFeatureUpgradeStatus
	// Begin upgrades for system features
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/migration-api-feature-upgrade.html
	PostFeatureUpgrade migration_post_feature_upgrade.NewPostFeatureUpgrade
}

type Profiling struct {
	// Extracts a UI-optimized structure to render flamegraphs from Universal
	// Profiling.
	// https://www.elastic.co/guide/en/observability/current/universal-profiling.html
	Flamegraph profiling_flamegraph.NewFlamegraph
	// Extracts raw stacktrace information from Universal Profiling.
	// https://www.elastic.co/guide/en/observability/current/universal-profiling.html
	Stacktraces profiling_stacktraces.NewStacktraces
	// Returns basic information about the status of Universal Profiling.
	// https://www.elastic.co/guide/en/observability/current/universal-profiling.html
	Status profiling_status.NewStatus
	// Extracts a list of topN functions from Universal Profiling.
	// https://www.elastic.co/guide/en/observability/current/universal-profiling.html
	TopnFunctions profiling_topn_functions.NewTopnFunctions
}

type QueryRules struct {
	// Deletes a query rule within a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-query-rule.html
	DeleteRule query_rules_delete_rule.NewDeleteRule
	// Deletes a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-query-ruleset.html
	DeleteRuleset query_rules_delete_ruleset.NewDeleteRuleset
	// Returns the details about a query rule within a query ruleset
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-query-rule.html
	GetRule query_rules_get_rule.NewGetRule
	// Returns the details about a query ruleset
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-query-ruleset.html
	GetRuleset query_rules_get_ruleset.NewGetRuleset
	// Returns summarized information about existing query rulesets.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-query-rulesets.html
	ListRulesets query_rules_list_rulesets.NewListRulesets
	// Creates or updates a query rule within a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-query-rule.html
	PutRule query_rules_put_rule.NewPutRule
	// Creates or updates a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-query-ruleset.html
	PutRuleset query_rules_put_ruleset.NewPutRuleset
	// Creates or updates a query ruleset.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/test-query-ruleset.html
	Test query_rules_test.NewTest
}

type Rollup struct {
	// Deletes an existing rollup job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-delete-job.html
	DeleteJob rollup_delete_job.NewDeleteJob
	// Retrieves the configuration, stats, and status of rollup jobs.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-get-job.html
	GetJobs rollup_get_jobs.NewGetJobs
	// Returns the capabilities of any rollup jobs that have been configured for a
	// specific index or index pattern.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-get-rollup-caps.html
	GetRollupCaps rollup_get_rollup_caps.NewGetRollupCaps
	// Creates a rollup job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-put-job.html
	PutJob rollup_put_job.NewPutJob
	// Enables searching rolled-up data using the standard Query DSL.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-search.html
	RollupSearch rollup_rollup_search.NewRollupSearch
	// Starts an existing, stopped rollup job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-start-job.html
	StartJob rollup_start_job.NewStartJob
	// Stops an existing, started rollup job.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-stop-job.html
	StopJob rollup_stop_job.NewStopJob
}

type SearchApplication struct {
	// Delete a search application.
	// Remove a search application and its associated alias. Indices attached to the
	// search application are not removed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-search-application.html
	Delete search_application_delete.NewDelete
	// Delete a behavioral analytics collection.
	// The associated data stream is also deleted.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-analytics-collection.html
	DeleteBehavioralAnalytics search_application_delete_behavioral_analytics.NewDeleteBehavioralAnalytics
	// Get search application details.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-search-application.html
	Get search_application_get.NewGet
	// Get behavioral analytics collections.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-analytics-collection.html
	GetBehavioralAnalytics search_application_get_behavioral_analytics.NewGetBehavioralAnalytics
	// Returns the existing search applications.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-search-applications.html
	List search_application_list.NewList
	// Create or update a search application.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-search-application.html
	Put search_application_put.NewPut
	// Create a behavioral analytics collection.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-analytics-collection.html
	PutBehavioralAnalytics search_application_put_behavioral_analytics.NewPutBehavioralAnalytics
	// Run a search application search.
	// Generate and run an Elasticsearch query that uses the specified query
	// parameteter and the search template associated with the search application or
	// default template.
	// Unspecified template parameters are assigned their default values if
	// applicable.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-application-search.html
	Search search_application_search.NewSearch
}

type Shutdown struct {
	// Removes a node from the shutdown list. Designed for indirect use by ECE/ESS
	// and ECK. Direct use is not supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current
	DeleteNode shutdown_delete_node.NewDeleteNode
	// Retrieve status of a node or nodes that are currently marked as shutting
	// down. Designed for indirect use by ECE/ESS and ECK. Direct use is not
	// supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current
	GetNode shutdown_get_node.NewGetNode
	// Adds a node to be shut down. Designed for indirect use by ECE/ESS and ECK.
	// Direct use is not supported.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current
	PutNode shutdown_put_node.NewPutNode
}

type Sql struct {
	// Clears the SQL cursor
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-sql-cursor-api.html
	ClearCursor sql_clear_cursor.NewClearCursor
	// Deletes an async SQL search or a stored synchronous SQL search. If the search
	// is still running, the API cancels it.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-async-sql-search-api.html
	DeleteAsync sql_delete_async.NewDeleteAsync
	// Returns the current status and available results for an async SQL search or
	// stored synchronous SQL search
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-sql-search-api.html
	GetAsync sql_get_async.NewGetAsync
	// Returns the current status of an async SQL search or a stored synchronous SQL
	// search
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-async-sql-search-status-api.html
	GetAsyncStatus sql_get_async_status.NewGetAsyncStatus
	// Executes a SQL request
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/sql-search-api.html
	Query sql_query.NewQuery
	// Translates SQL into Elasticsearch queries
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/sql-translate-api.html
	Translate sql_translate.NewTranslate
}

type Ssl struct {
	// Get SSL certificates.
	//
	// Get information about the X.509 certificates that are used to encrypt
	// communications in the cluster.
	// The API returns a list that includes certificates from all TLS contexts
	// including:
	//
	// - Settings for transport and HTTP interfaces
	// - TLS settings that are used within authentication realms
	// - TLS settings for remote monitoring exporters
	//
	// The list includes certificates that are used for configuring trust, such as
	// those configured in the `xpack.security.transport.ssl.truststore` and
	// `xpack.security.transport.ssl.certificate_authorities` settings.
	// It also includes certificates that are used for configuring server identity,
	// such as `xpack.security.http.ssl.keystore` and
	// `xpack.security.http.ssl.certificate settings`.
	//
	// The list does not include certificates that are sourced from the default SSL
	// context of the Java Runtime Environment (JRE), even if those certificates are
	// in use within Elasticsearch.
	//
	// NOTE: When a PKCS#11 token is configured as the truststore of the JRE, the
	// API returns all the certificates that are included in the PKCS#11 token
	// irrespective of whether these are used in the Elasticsearch TLS
	// configuration.
	//
	// If Elasticsearch is configured to use a keystore or truststore, the API
	// output includes all certificates in that store, even though some of the
	// certificates might not be in active use within the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-ssl.html
	Certificates ssl_certificates.NewCertificates
}

type Synonyms struct {
	// Deletes a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-synonyms-set.html
	DeleteSynonym synonyms_delete_synonym.NewDeleteSynonym
	// Deletes a synonym rule in a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-synonym-rule.html
	DeleteSynonymRule synonyms_delete_synonym_rule.NewDeleteSynonymRule
	// Retrieves a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-synonyms-set.html
	GetSynonym synonyms_get_synonym.NewGetSynonym
	// Retrieves a synonym rule from a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-synonym-rule.html
	GetSynonymRule synonyms_get_synonym_rule.NewGetSynonymRule
	// Retrieves a summary of all defined synonym sets
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/list-synonyms-sets.html
	GetSynonymsSets synonyms_get_synonyms_sets.NewGetSynonymsSets
	// Creates or updates a synonym set.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-synonyms-set.html
	PutSynonym synonyms_put_synonym.NewPutSynonym
	// Creates or updates a synonym rule in a synonym set
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-synonym-rule.html
	PutSynonymRule synonyms_put_synonym_rule.NewPutSynonymRule
}

type TextStructure struct {
	// Finds the structure of a text field in an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-field-structure.html
	FindFieldStructure text_structure_find_field_structure.NewFindFieldStructure
	// Finds the structure of a list of messages. The messages must contain data
	// that is suitable to be ingested into Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-message-structure.html
	FindMessageStructure text_structure_find_message_structure.NewFindMessageStructure
	// Finds the structure of a text file. The text file must contain data that is
	// suitable to be ingested into Elasticsearch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/find-structure.html
	FindStructure text_structure_find_structure.NewFindStructure
	// Tests a Grok pattern on some text.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/test-grok-pattern.html
	TestGrokPattern text_structure_test_grok_pattern.NewTestGrokPattern
}

type Transform struct {
	// Delete a transform.
	// Deletes a transform.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/delete-transform.html
	DeleteTransform transform_delete_transform.NewDeleteTransform
	// Retrieves transform usage information for transform nodes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform-node-stats.html
	GetNodeStats transform_get_node_stats.NewGetNodeStats
	// Get transforms.
	// Retrieves configuration information for transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform.html
	GetTransform transform_get_transform.NewGetTransform
	// Get transform stats.
	// Retrieves usage information for transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform-stats.html
	GetTransformStats transform_get_transform_stats.NewGetTransformStats
	// Preview a transform.
	// Generates a preview of the results that you will get when you create a
	// transform with the same configuration.
	//
	// It returns a maximum of 100 results. The calculations are based on all the
	// current data in the source index. It also
	// generates a list of mappings and settings for the destination index. These
	// values are determined based on the field
	// types of the source index and the transform aggregations.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/preview-transform.html
	PreviewTransform transform_preview_transform.NewPreviewTransform
	// Create a transform.
	// Creates a transform.
	//
	// A transform copies data from source indices, transforms it, and persists it
	// into an entity-centric destination index. You can also think of the
	// destination index as a two-dimensional tabular data structure (known as
	// a data frame). The ID for each document in the data frame is generated from a
	// hash of the entity, so there is a
	// unique row per entity.
	//
	// You must choose either the latest or pivot method for your transform; you
	// cannot use both in a single transform. If
	// you choose to use the pivot method for your transform, the entities are
	// defined by the set of `group_by` fields in
	// the pivot object. If you choose to use the latest method, the entities are
	// defined by the `unique_key` field values
	// in the latest object.
	//
	// You must have `create_index`, `index`, and `read` privileges on the
	// destination index and `read` and
	// `view_index_metadata` privileges on the source indices. When Elasticsearch
	// security features are enabled, the
	// transform remembers which roles the user that created it had at the time of
	// creation and uses those same roles. If
	// those roles do not have the required privileges on the source and destination
	// indices, the transform fails when it
	// attempts unauthorized operations.
	//
	// NOTE: You must use Kibana or this API to create a transform. Do not add a
	// transform directly into any
	// `.transform-internal*` indices using the Elasticsearch index API. If
	// Elasticsearch security features are enabled, do
	// not give users any privileges on `.transform-internal*` indices. If you used
	// transforms prior to 7.5, also do not
	// give users any privileges on `.data-frame-internal*` indices.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-transform.html
	PutTransform transform_put_transform.NewPutTransform
	// Reset a transform.
	// Resets a transform.
	// Before you can reset it, you must stop it; alternatively, use the `force`
	// query parameter.
	// If the destination index was created by the transform, it is deleted.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/reset-transform.html
	ResetTransform transform_reset_transform.NewResetTransform
	// Schedule a transform to start now.
	// Instantly runs a transform to process data.
	//
	// If you _schedule_now a transform, it will process the new data instantly,
	// without waiting for the configured frequency interval. After _schedule_now
	// API is called,
	// the transform will be processed again at now + frequency unless _schedule_now
	// API
	// is called again in the meantime.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/schedule-now-transform.html
	ScheduleNowTransform transform_schedule_now_transform.NewScheduleNowTransform
	// Start a transform.
	// Starts a transform.
	//
	// When you start a transform, it creates the destination index if it does not
	// already exist. The `number_of_shards` is
	// set to `1` and the `auto_expand_replicas` is set to `0-1`. If it is a pivot
	// transform, it deduces the mapping
	// definitions for the destination index from the source indices and the
	// transform aggregations. If fields in the
	// destination index are derived from scripts (as in the case of
	// `scripted_metric` or `bucket_script` aggregations),
	// the transform uses dynamic mappings unless an index template exists. If it is
	// a latest transform, it does not deduce
	// mapping definitions; it uses dynamic mappings. To use explicit mappings,
	// create the destination index before you
	// start the transform. Alternatively, you can create an index template, though
	// it does not affect the deduced mappings
	// in a pivot transform.
	//
	// When the transform starts, a series of validations occur to ensure its
	// success. If you deferred validation when you
	// created the transform, they occur when you start the transform—​with the
	// exception of privilege checks. When
	// Elasticsearch security features are enabled, the transform remembers which
	// roles the user that created it had at the
	// time of creation and uses those same roles. If those roles do not have the
	// required privileges on the source and
	// destination indices, the transform fails when it attempts unauthorized
	// operations.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/start-transform.html
	StartTransform transform_start_transform.NewStartTransform
	// Stop transforms.
	// Stops one or more transforms.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/stop-transform.html
	StopTransform transform_stop_transform.NewStopTransform
	// Update a transform.
	// Updates certain properties of a transform.
	//
	// All updated properties except `description` do not take effect until after
	// the transform starts the next checkpoint,
	// thus there is data consistency in each checkpoint. To use this API, you must
	// have `read` and `view_index_metadata`
	// privileges for the source indices. You must also have `index` and `read`
	// privileges for the destination index. When
	// Elasticsearch security features are enabled, the transform remembers which
	// roles the user who updated it had at the
	// time of update and runs with those privileges.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/update-transform.html
	UpdateTransform transform_update_transform.NewUpdateTransform
	// Upgrades all transforms.
	// This API identifies transforms that have a legacy configuration format and
	// upgrades them to the latest version. It
	// also cleans up the internal data structures that store the transform state
	// and checkpoints. The upgrade does not
	// affect the source and destination indices. The upgrade also does not affect
	// the roles that transforms use when
	// Elasticsearch security features are enabled; the role used to read source
	// data and write to the destination index
	// remains unchanged.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/upgrade-transforms.html
	UpgradeTransforms transform_upgrade_transforms.NewUpgradeTransforms
}

type Watcher struct {
	// Acknowledges a watch, manually throttling the execution of the watch's
	// actions.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-ack-watch.html
	AckWatch watcher_ack_watch.NewAckWatch
	// Activates a currently inactive watch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-activate-watch.html
	ActivateWatch watcher_activate_watch.NewActivateWatch
	// Deactivates a currently active watch.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-deactivate-watch.html
	DeactivateWatch watcher_deactivate_watch.NewDeactivateWatch
	// Removes a watch from Watcher.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-delete-watch.html
	DeleteWatch watcher_delete_watch.NewDeleteWatch
	// This API can be used to force execution of the watch outside of its
	// triggering logic or to simulate the watch execution for debugging purposes.
	// For testing and debugging purposes, you also have fine-grained control on how
	// the watch runs. You can execute the watch without executing all of its
	// actions or alternatively by simulating them. You can also force execution by
	// ignoring the watch condition and control whether a watch record would be
	// written to the watch history after execution.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-execute-watch.html
	ExecuteWatch watcher_execute_watch.NewExecuteWatch
	// Retrieve settings for the watcher system index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-get-settings.html
	GetSettings watcher_get_settings.NewGetSettings
	// Retrieves a watch by its ID.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-get-watch.html
	GetWatch watcher_get_watch.NewGetWatch
	// Creates a new watch, or updates an existing one.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-put-watch.html
	PutWatch watcher_put_watch.NewPutWatch
	// Retrieves stored watches.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-query-watches.html
	QueryWatches watcher_query_watches.NewQueryWatches
	// Starts Watcher if it is not already running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-start.html
	Start watcher_start.NewStart
	// Retrieves the current Watcher metrics.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-stats.html
	Stats watcher_stats.NewStats
	// Stops Watcher if it is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-stop.html
	Stop watcher_stop.NewStop
	// Update settings for the watcher system index
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-update-settings.html
	UpdateSettings watcher_update_settings.NewUpdateSettings
}

type API struct {
	AsyncSearch         AsyncSearch
	Autoscaling         Autoscaling
	Capabilities        Capabilities
	Ccr                 Ccr
	Connector           Connector
	Core                Core
	DanglingIndices     DanglingIndices
	Enrich              Enrich
	Eql                 Eql
	Esql                Esql
	Features            Features
	Fleet               Fleet
	Graph               Graph
	Ilm                 Ilm
	Inference           Inference
	Ingest              Ingest
	License             License
	Logstash            Logstash
	Migration           Migration
	Profiling           Profiling
	QueryRules          QueryRules
	Rollup              Rollup
	SearchApplication   SearchApplication
	Shutdown            Shutdown
	Sql                 Sql
	Ssl                 Ssl
	Synonyms            Synonyms
	TextStructure       TextStructure
	Transform           Transform
	Watcher             Watcher

	// Bulk index or delete documents.
	// Performs multiple indexing or delete operations in a single API call.
	// This reduces overhead and can greatly increase indexing speed.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
	Bulk core_bulk.NewBulk
	// Clear a scrolling search.
	//
	// Clear the search context and results for a scrolling search.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/clear-scroll-api.html
	ClearScroll core_clear_scroll.NewClearScroll
	// Close a point in time.
	//
	// A point in time must be opened explicitly before being used in search
	// requests.
	// The `keep_alive` parameter tells Elasticsearch how long it should persist.
	// A point in time is automatically closed when the `keep_alive` period has
	// elapsed.
	// However, keeping points in time has a cost; close them as soon as they are no
	// longer required for search requests.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	ClosePointInTime core_close_point_in_time.NewClosePointInTime
	// Returns number of documents matching a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-count.html
	Count core_count.NewCount
	// Index a document.
	// Adds a JSON document to the specified data stream or index and makes it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Create core_create.NewCreate
	// Delete a document.
	// Removes a JSON document from the specified index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete.html
	Delete core_delete.NewDelete
	// Delete documents.
	// Deletes documents that match the specified query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQuery core_delete_by_query.NewDeleteByQuery
	// Throttle a delete by query operation.
	//
	// Change the number of requests per second for a particular delete by query
	// operation.
	// Rethrottling that speeds up the query takes effect immediately but
	// rethrotting that slows down the query takes effect after completing the
	// current batch to prevent scroll timeouts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
	DeleteByQueryRethrottle core_delete_by_query_rethrottle.NewDeleteByQueryRethrottle
	// Delete a script or search template.
	// Deletes a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	DeleteScript core_delete_script.NewDeleteScript
	// Check a document.
	// Checks if a specified document exists.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Exists core_exists.NewExists
	// Check for a document source.
	// Checks if a document's `_source` is stored.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	ExistsSource core_exists_source.NewExistsSource
	// Explain a document match result.
	// Returns information about why a specific document matches, or doesn’t match,
	// a query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-explain.html
	Explain core_explain.NewExplain
	// Get the field capabilities.
	//
	// Get information about the capabilities of fields among multiple indices.
	//
	// For data streams, the API returns field capabilities among the stream’s
	// backing indices.
	// It returns runtime fields like any other field.
	// For example, a runtime field with a type of keyword is returned the same as
	// any other field that belongs to the `keyword` family.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-field-caps.html
	FieldCaps core_field_caps.NewFieldCaps
	// Get a document by its ID.
	// Retrieves the document with the specified ID from an index.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	Get core_get.NewGet
	// Get a script or search template.
	// Retrieves a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	GetScript core_get_script.NewGetScript
	// Get script contexts.
	//
	// Get a list of supported script contexts and their methods.
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-contexts.html
	GetScriptContext core_get_script_context.NewGetScriptContext
	// Get script languages.
	//
	// Get a list of available script types, languages, and contexts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	GetScriptLanguages core_get_script_languages.NewGetScriptLanguages
	// Get a document's source.
	// Returns the source of a document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html
	GetSource core_get_source.NewGetSource
	// Returns the health of the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/health-api.html
	HealthReport core_health_report.NewHealthReport
	// Index a document.
	// Adds a JSON document to the specified data stream or index and makes it
	// searchable.
	// If the target is an index and the document already exists, the request
	// updates the document and increments its version.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	Index core_index.NewIndex
	// Get cluster info.
	// Returns basic information about the cluster.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Info core_info.NewInfo
	// Run a knn search.
	//
	// NOTE: The kNN search API has been replaced by the `knn` option in the search
	// API.
	//
	// Perform a k-nearest neighbor (kNN) search on a dense_vector field and return
	// the matching documents.
	// Given a query vector, the API finds the k closest vectors and returns those
	// documents as search hits.
	//
	// Elasticsearch uses the HNSW algorithm to support efficient kNN search.
	// Like most kNN algorithms, HNSW is an approximate method that sacrifices
	// result accuracy for improved search speed.
	// This means the results returned are not always the true k closest neighbors.
	//
	// The kNN search API supports restricting the search using a filter.
	// The search will return the top k documents that also match the filter query.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	KnnSearch core_knn_search.NewKnnSearch
	// Get multiple documents.
	//
	// Get multiple JSON documents by ID from one or more indices.
	// If you specify an index in the request URI, you only need to specify the
	// document IDs in the request body.
	// To ensure fast responses, this multi get (mget) API responds with partial
	// results if one or more shards fail.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-get.html
	Mget core_mget.NewMget
	// Run multiple searches.
	//
	// The format of the request is similar to the bulk API format and makes use of
	// the newline delimited JSON (NDJSON) format.
	// The structure is as follows:
	//
	// ```
	// header\n
	// body\n
	// header\n
	// body\n
	// ```
	//
	// This structure is specifically optimized to reduce parsing if a specific
	// search ends up redirected to another node.
	//
	// IMPORTANT: The final line of data must end with a newline character `\n`.
	// Each newline character may be preceded by a carriage return `\r`.
	// When sending requests to this endpoint the `Content-Type` header should be
	// set to `application/x-ndjson`.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	Msearch core_msearch.NewMsearch
	// Run multiple templated searches.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-multi-search.html
	MsearchTemplate core_msearch_template.NewMsearchTemplate
	// Get multiple term vectors.
	//
	// You can specify existing documents by index and ID or provide artificial
	// documents in the body of the request.
	// You can specify the index in the request body or request URI.
	// The response contains a `docs` array with all the fetched termvectors.
	// Each element has the structure provided by the termvectors API.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-multi-termvectors.html
	Mtermvectors core_mtermvectors.NewMtermvectors
	// Open a point in time.
	//
	// A search request by default runs against the most recent visible data of the
	// target indices,
	// which is called point in time. Elasticsearch pit (point in time) is a
	// lightweight view into the
	// state of the data as it existed when initiated. In some cases, it’s preferred
	// to perform multiple
	// search requests using the same point in time. For example, if refreshes
	// happen between
	// `search_after` requests, then the results of those requests might not be
	// consistent as changes happening
	// between searches are only visible to the more recent point in time.
	//
	// A point in time must be opened explicitly before being used in search
	// requests.
	// The `keep_alive` parameter tells Elasticsearch how long it should persist.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
	OpenPointInTime core_open_point_in_time.NewOpenPointInTime
	// Ping the cluster.
	// Returns whether the cluster is running.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
	Ping core_ping.NewPing
	// Create or update a script or search template.
	// Creates or updates a stored script or search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html
	PutScript core_put_script.NewPutScript
	// Evaluate ranked search results.
	//
	// Evaluate the quality of ranked search results over a set of typical search
	// queries.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-rank-eval.html
	RankEval core_rank_eval.NewRankEval
	// Reindex documents.
	// Copies documents from a source to a destination. The source can be any
	// existing index, alias, or data stream. The destination must differ from the
	// source. For example, you cannot reindex a data stream into itself.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	Reindex core_reindex.NewReindex
	// Throttle a reindex operation.
	//
	// Change the number of requests per second for a particular reindex operation.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-reindex.html
	ReindexRethrottle core_reindex_rethrottle.NewReindexRethrottle
	// Render a search template.
	//
	// Render a search template as a search request body.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/render-search-template-api.html
	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate
	// Run a script.
	// Runs a script and returns a result.
	// https://www.elastic.co/guide/en/elasticsearch/painless/current/painless-execute-api.html
	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute
	// Run a scrolling search.
	//
	// IMPORTANT: The scroll API is no longer recommend for deep pagination. If you
	// need to preserve the index state while paging through more than 10,000 hits,
	// use the `search_after` parameter with a point in time (PIT).
	//
	// The scroll API gets large sets of results from a single scrolling search
	// request.
	// To get the necessary scroll ID, submit a search API request that includes an
	// argument for the `scroll` query parameter.
	// The `scroll` parameter indicates how long Elasticsearch should retain the
	// search context for the request.
	// The search response returns a scroll ID in the `_scroll_id` response body
	// parameter.
	// You can then use the scroll ID with the scroll API to retrieve the next batch
	// of results for the request.
	// If the Elasticsearch security features are enabled, the access to the results
	// of a specific scroll ID is restricted to the user or API key that submitted
	// the search.
	//
	// You can also use the scroll API to specify a new scroll parameter that
	// extends or shortens the retention period for the search context.
	//
	// IMPORTANT: Results from a scrolling search reflect the state of the index at
	// the time of the initial search request. Subsequent indexing or document
	// changes only affect later search and scroll requests.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-body.html#request-body-search-scroll
	Scroll core_scroll.NewScroll
	// Run a search.
	//
	// Get search hits that match the query defined in the request.
	// You can provide search queries using the `q` query string parameter or the
	// request body.
	// If both are specified, only the query parameter is used.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
	Search core_search.NewSearch
	// Search a vector tile.
	//
	// Search a vector tile for geospatial values.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-vector-tile-api.html
	SearchMvt core_search_mvt.NewSearchMvt
	// Get the search shards.
	//
	// Get the indices and shards that a search request would be run against.
	// This information can be useful for working out issues or planning
	// optimizations with routing and shard preferences.
	// When filtered aliases are used, the filter is returned as part of the indices
	// section.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-shards.html
	SearchShards core_search_shards.NewSearchShards
	// Run a search with a search template.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template.html
	SearchTemplate core_search_template.NewSearchTemplate
	// Get terms in an index.
	//
	// Discover terms that match a partial string in an index.
	// This "terms enum" API is designed for low-latency look-ups used in
	// auto-complete scenarios.
	//
	// If the `complete` property in the response is false, the returned terms set
	// may be incomplete and should be treated as approximate.
	// This can occur due to a few reasons, such as a request timeout or a node
	// error.
	//
	// NOTE: The terms enum API may return terms from deleted documents. Deleted
	// documents are initially only marked as deleted. It is not until their
	// segments are merged that documents are actually deleted. Until that happens,
	// the terms enum API will return terms from these documents.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-terms-enum.html
	TermsEnum core_terms_enum.NewTermsEnum
	// Get term vector information.
	//
	// Get information and statistics about terms in the fields of a particular
	// document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-termvectors.html
	Termvectors core_termvectors.NewTermvectors
	// Update a document.
	// Updates a document by running a script or passing a partial document.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update.html
	Update core_update.NewUpdate
	// Update documents.
	// Updates documents that match the specified query.
	// If no query is specified, performs an update on every document in the data
	// stream or index without modifying the source, which is useful for picking up
	// mapping changes.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
	UpdateByQuery core_update_by_query.NewUpdateByQuery
	// Throttle an update by query operation.
	//
	// Change the number of requests per second for a particular update by query
	// operation.
	// Rethrottling that speeds up the query takes effect immediately but
	// rethrotting that slows down the query takes effect after completing the
	// current batch to prevent scroll timeouts.
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update-by-query.html
	UpdateByQueryRethrottle core_update_by_query_rethrottle.NewUpdateByQueryRethrottle
}

func New(tp elastictransport.Interface) *API {
	return &API{
		// AsyncSearch
		AsyncSearch: AsyncSearch{
			Delete: async_search_delete.NewDeleteFunc(tp),
			Get:    async_search_get.NewGetFunc(tp),
			Status: async_search_status.NewStatusFunc(tp),
			Submit: async_search_submit.NewSubmitFunc(tp),
		},

		// Autoscaling
		Autoscaling: Autoscaling{
			DeleteAutoscalingPolicy: autoscaling_delete_autoscaling_policy.NewDeleteAutoscalingPolicyFunc(tp),
			GetAutoscalingCapacity:  autoscaling_get_autoscaling_capacity.NewGetAutoscalingCapacityFunc(tp),
			GetAutoscalingPolicy:    autoscaling_get_autoscaling_policy.NewGetAutoscalingPolicyFunc(tp),
			PutAutoscalingPolicy:    autoscaling_put_autoscaling_policy.NewPutAutoscalingPolicyFunc(tp),
		},

		// Capabilities
		Capabilities: Capabilities{
			Capabilities: capabilities.NewCapabilitiesFunc(tp),
		},

		// Ccr
		Ccr: Ccr{
			DeleteAutoFollowPattern: ccr_delete_auto_follow_pattern.NewDeleteAutoFollowPatternFunc(tp),
			Follow:                  ccr_follow.NewFollowFunc(tp),
			FollowInfo:              ccr_follow_info.NewFollowInfoFunc(tp),
			FollowStats:             ccr_follow_stats.NewFollowStatsFunc(tp),
			ForgetFollower:          ccr_forget_follower.NewForgetFollowerFunc(tp),
			GetAutoFollowPattern:    ccr_get_auto_follow_pattern.NewGetAutoFollowPatternFunc(tp),
			PauseAutoFollowPattern:  ccr_pause_auto_follow_pattern.NewPauseAutoFollowPatternFunc(tp),
			PauseFollow:             ccr_pause_follow.NewPauseFollowFunc(tp),
			PutAutoFollowPattern:    ccr_put_auto_follow_pattern.NewPutAutoFollowPatternFunc(tp),
			ResumeAutoFollowPattern: ccr_resume_auto_follow_pattern.NewResumeAutoFollowPatternFunc(tp),
			ResumeFollow:            ccr_resume_follow.NewResumeFollowFunc(tp),
			Stats:                   ccr_stats.NewStatsFunc(tp),
			Unfollow:                ccr_unfollow.NewUnfollowFunc(tp),
		},

		// Connector
		Connector: Connector{
			CheckIn:                   connector_check_in.NewCheckInFunc(tp),
			Delete:                    connector_delete.NewDeleteFunc(tp),
			Get:                       connector_get.NewGetFunc(tp),
			LastSync:                  connector_last_sync.NewLastSyncFunc(tp),
			List:                      connector_list.NewListFunc(tp),
			Post:                      connector_post.NewPostFunc(tp),
			Put:                       connector_put.NewPutFunc(tp),
			SecretPost:                connector_secret_post.NewSecretPostFunc(tp),
			SyncJobCancel:             connector_sync_job_cancel.NewSyncJobCancelFunc(tp),
			SyncJobDelete:             connector_sync_job_delete.NewSyncJobDeleteFunc(tp),
			SyncJobGet:                connector_sync_job_get.NewSyncJobGetFunc(tp),
			SyncJobList:               connector_sync_job_list.NewSyncJobListFunc(tp),
			SyncJobPost:               connector_sync_job_post.NewSyncJobPostFunc(tp),
			UpdateActiveFiltering:     connector_update_active_filtering.NewUpdateActiveFilteringFunc(tp),
			UpdateApiKeyId:            connector_update_api_key_id.NewUpdateApiKeyIdFunc(tp),
			UpdateConfiguration:       connector_update_configuration.NewUpdateConfigurationFunc(tp),
			UpdateError:               connector_update_error.NewUpdateErrorFunc(tp),
			UpdateFiltering:           connector_update_filtering.NewUpdateFilteringFunc(tp),
			UpdateFilteringValidation: connector_update_filtering_validation.NewUpdateFilteringValidationFunc(tp),
			UpdateIndexName:           connector_update_index_name.NewUpdateIndexNameFunc(tp),
			UpdateName:                connector_update_name.NewUpdateNameFunc(tp),
			UpdateNative:              connector_update_native.NewUpdateNativeFunc(tp),
			UpdatePipeline:            connector_update_pipeline.NewUpdatePipelineFunc(tp),
			UpdateScheduling:          connector_update_scheduling.NewUpdateSchedulingFunc(tp),
			UpdateServiceType:         connector_update_service_type.NewUpdateServiceTypeFunc(tp),
			UpdateStatus:              connector_update_status.NewUpdateStatusFunc(tp),
		},

		// Core
		Core: Core{
			Bulk:                    core_bulk.NewBulkFunc(tp),
			ClearScroll:             core_clear_scroll.NewClearScrollFunc(tp),
			ClosePointInTime:        core_close_point_in_time.NewClosePointInTimeFunc(tp),
			Count:                   core_count.NewCountFunc(tp),
			Create:                  core_create.NewCreateFunc(tp),
			Delete:                  core_delete.NewDeleteFunc(tp),
			DeleteByQuery:           core_delete_by_query.NewDeleteByQueryFunc(tp),
			DeleteByQueryRethrottle: core_delete_by_query_rethrottle.NewDeleteByQueryRethrottleFunc(tp),
			DeleteScript:            core_delete_script.NewDeleteScriptFunc(tp),
			Exists:                  core_exists.NewExistsFunc(tp),
			ExistsSource:            core_exists_source.NewExistsSourceFunc(tp),
			Explain:                 core_explain.NewExplainFunc(tp),
			FieldCaps:               core_field_caps.NewFieldCapsFunc(tp),
			Get:                     core_get.NewGetFunc(tp),
			GetScript:               core_get_script.NewGetScriptFunc(tp),
			GetScriptContext:        core_get_script_context.NewGetScriptContextFunc(tp),
			GetScriptLanguages:      core_get_script_languages.NewGetScriptLanguagesFunc(tp),
			GetSource:               core_get_source.NewGetSourceFunc(tp),
			HealthReport:            core_health_report.NewHealthReportFunc(tp),
			Index:                   core_index.NewIndexFunc(tp),
			Info:                    core_info.NewInfoFunc(tp),
			KnnSearch:               core_knn_search.NewKnnSearchFunc(tp),
			Mget:                    core_mget.NewMgetFunc(tp),
			Msearch:                 core_msearch.NewMsearchFunc(tp),
			MsearchTemplate:         core_msearch_template.NewMsearchTemplateFunc(tp),
			Mtermvectors:            core_mtermvectors.NewMtermvectorsFunc(tp),
			OpenPointInTime:         core_open_point_in_time.NewOpenPointInTimeFunc(tp),
			Ping:                    core_ping.NewPingFunc(tp),
			PutScript:               core_put_script.NewPutScriptFunc(tp),
			RankEval:                core_rank_eval.NewRankEvalFunc(tp),
			Reindex:                 core_reindex.NewReindexFunc(tp),
			ReindexRethrottle:       core_reindex_rethrottle.NewReindexRethrottleFunc(tp),
			RenderSearchTemplate:    core_render_search_template.NewRenderSearchTemplateFunc(tp),
			ScriptsPainlessExecute:  core_scripts_painless_execute.NewScriptsPainlessExecuteFunc(tp),
			Scroll:                  core_scroll.NewScrollFunc(tp),
			Search:                  core_search.NewSearchFunc(tp),
			SearchMvt:               core_search_mvt.NewSearchMvtFunc(tp),
			SearchShards:            core_search_shards.NewSearchShardsFunc(tp),
			SearchTemplate:          core_search_template.NewSearchTemplateFunc(tp),
			TermsEnum:               core_terms_enum.NewTermsEnumFunc(tp),
			Termvectors:             core_termvectors.NewTermvectorsFunc(tp),
			Update:                  core_update.NewUpdateFunc(tp),
			UpdateByQuery:           core_update_by_query.NewUpdateByQueryFunc(tp),
			UpdateByQueryRethrottle: core_update_by_query_rethrottle.NewUpdateByQueryRethrottleFunc(tp),
		},

		// DanglingIndices
		DanglingIndices: DanglingIndices{
			DeleteDanglingIndex: dangling_indices_delete_dangling_index.NewDeleteDanglingIndexFunc(tp),
			ImportDanglingIndex: dangling_indices_import_dangling_index.NewImportDanglingIndexFunc(tp),
			ListDanglingIndices: dangling_indices_list_dangling_indices.NewListDanglingIndicesFunc(tp),
		},

		// Enrich
		Enrich: Enrich{
			DeletePolicy:  enrich_delete_policy.NewDeletePolicyFunc(tp),
			ExecutePolicy: enrich_execute_policy.NewExecutePolicyFunc(tp),
			GetPolicy:     enrich_get_policy.NewGetPolicyFunc(tp),
			PutPolicy:     enrich_put_policy.NewPutPolicyFunc(tp),
			Stats:         enrich_stats.NewStatsFunc(tp),
		},

		// Eql
		Eql: Eql{
			Delete:    eql_delete.NewDeleteFunc(tp),
			Get:       eql_get.NewGetFunc(tp),
			GetStatus: eql_get_status.NewGetStatusFunc(tp),
			Search:    eql_search.NewSearchFunc(tp),
		},

		// Esql
		Esql: Esql{
			AsyncQuery: esql_async_query.NewAsyncQueryFunc(tp),
			Query:      esql_query.NewQueryFunc(tp),
		},

		// Features
		Features: Features{
			GetFeatures:   features_get_features.NewGetFeaturesFunc(tp),
			ResetFeatures: features_reset_features.NewResetFeaturesFunc(tp),
		},

		// Fleet
		Fleet: Fleet{
			GlobalCheckpoints: fleet_global_checkpoints.NewGlobalCheckpointsFunc(tp),
			Msearch:           fleet_msearch.NewMsearchFunc(tp),
			PostSecret:        fleet_post_secret.NewPostSecretFunc(tp),
			Search:            fleet_search.NewSearchFunc(tp),
		},

		// Graph
		Graph: Graph{
			Explore: graph_explore.NewExploreFunc(tp),
		},

		// Ilm
		Ilm: Ilm{
			DeleteLifecycle:    ilm_delete_lifecycle.NewDeleteLifecycleFunc(tp),
			ExplainLifecycle:   ilm_explain_lifecycle.NewExplainLifecycleFunc(tp),
			GetLifecycle:       ilm_get_lifecycle.NewGetLifecycleFunc(tp),
			GetStatus:          ilm_get_status.NewGetStatusFunc(tp),
			MigrateToDataTiers: ilm_migrate_to_data_tiers.NewMigrateToDataTiersFunc(tp),
			MoveToStep:         ilm_move_to_step.NewMoveToStepFunc(tp),
			PutLifecycle:       ilm_put_lifecycle.NewPutLifecycleFunc(tp),
			RemovePolicy:       ilm_remove_policy.NewRemovePolicyFunc(tp),
			Retry:              ilm_retry.NewRetryFunc(tp),
			Start:              ilm_start.NewStartFunc(tp),
			Stop:               ilm_stop.NewStopFunc(tp),
		},

		// Inference
		Inference: Inference{
			Delete:    inference_delete.NewDeleteFunc(tp),
			Get:       inference_get.NewGetFunc(tp),
			Inference: inference_inference.NewInferenceFunc(tp),
			Put:       inference_put.NewPutFunc(tp),
		},

		// Ingest
		Ingest: Ingest{
			DeleteGeoipDatabase: ingest_delete_geoip_database.NewDeleteGeoipDatabaseFunc(tp),
			DeletePipeline:      ingest_delete_pipeline.NewDeletePipelineFunc(tp),
			GeoIpStats:          ingest_geo_ip_stats.NewGeoIpStatsFunc(tp),
			GetGeoipDatabase:    ingest_get_geoip_database.NewGetGeoipDatabaseFunc(tp),
			GetPipeline:         ingest_get_pipeline.NewGetPipelineFunc(tp),
			ProcessorGrok:       ingest_processor_grok.NewProcessorGrokFunc(tp),
			PutGeoipDatabase:    ingest_put_geoip_database.NewPutGeoipDatabaseFunc(tp),
			PutPipeline:         ingest_put_pipeline.NewPutPipelineFunc(tp),
			Simulate:            ingest_simulate.NewSimulateFunc(tp),
		},

		// License
		License: License{
			Delete:         license_delete.NewDeleteFunc(tp),
			Get:            license_get.NewGetFunc(tp),
			GetBasicStatus: license_get_basic_status.NewGetBasicStatusFunc(tp),
			GetTrialStatus: license_get_trial_status.NewGetTrialStatusFunc(tp),
			Post:           license_post.NewPostFunc(tp),
			PostStartBasic: license_post_start_basic.NewPostStartBasicFunc(tp),
			PostStartTrial: license_post_start_trial.NewPostStartTrialFunc(tp),
		},

		// Logstash
		Logstash: Logstash{
			DeletePipeline: logstash_delete_pipeline.NewDeletePipelineFunc(tp),
			GetPipeline:    logstash_get_pipeline.NewGetPipelineFunc(tp),
			PutPipeline:    logstash_put_pipeline.NewPutPipelineFunc(tp),
		},

		// Migration
		Migration: Migration{
			Deprecations:            migration_deprecations.NewDeprecationsFunc(tp),
			GetFeatureUpgradeStatus: migration_get_feature_upgrade_status.NewGetFeatureUpgradeStatusFunc(tp),
			PostFeatureUpgrade:      migration_post_feature_upgrade.NewPostFeatureUpgradeFunc(tp),
		},

		// Profiling
		Profiling: Profiling{
			Flamegraph:    profiling_flamegraph.NewFlamegraphFunc(tp),
			Stacktraces:   profiling_stacktraces.NewStacktracesFunc(tp),
			Status:        profiling_status.NewStatusFunc(tp),
			TopnFunctions: profiling_topn_functions.NewTopnFunctionsFunc(tp),
		},

		// QueryRules
		QueryRules: QueryRules{
			DeleteRule:    query_rules_delete_rule.NewDeleteRuleFunc(tp),
			DeleteRuleset: query_rules_delete_ruleset.NewDeleteRulesetFunc(tp),
			GetRule:       query_rules_get_rule.NewGetRuleFunc(tp),
			GetRuleset:    query_rules_get_ruleset.NewGetRulesetFunc(tp),
			ListRulesets:  query_rules_list_rulesets.NewListRulesetsFunc(tp),
			PutRule:       query_rules_put_rule.NewPutRuleFunc(tp),
			PutRuleset:    query_rules_put_ruleset.NewPutRulesetFunc(tp),
			Test:          query_rules_test.NewTestFunc(tp),
		},

		// Rollup
		Rollup: Rollup{
			DeleteJob:          rollup_delete_job.NewDeleteJobFunc(tp),
			GetJobs:            rollup_get_jobs.NewGetJobsFunc(tp),
			GetRollupCaps:      rollup_get_rollup_caps.NewGetRollupCapsFunc(tp),
			PutJob:             rollup_put_job.NewPutJobFunc(tp),
			RollupSearch:       rollup_rollup_search.NewRollupSearchFunc(tp),
			StartJob:           rollup_start_job.NewStartJobFunc(tp),
			StopJob:            rollup_stop_job.NewStopJobFunc(tp),
		},

		// SearchApplication
		SearchApplication: SearchApplication{
			Delete:                    search_application_delete.NewDeleteFunc(tp),
			DeleteBehavioralAnalytics: search_application_delete_behavioral_analytics.NewDeleteBehavioralAnalyticsFunc(tp),
			Get:                       search_application_get.NewGetFunc(tp),
			GetBehavioralAnalytics:    search_application_get_behavioral_analytics.NewGetBehavioralAnalyticsFunc(tp),
			List:                      search_application_list.NewListFunc(tp),
			Put:                       search_application_put.NewPutFunc(tp),
			PutBehavioralAnalytics:    search_application_put_behavioral_analytics.NewPutBehavioralAnalyticsFunc(tp),
			Search:                    search_application_search.NewSearchFunc(tp),
		},

		// Shutdown
		Shutdown: Shutdown{
			DeleteNode: shutdown_delete_node.NewDeleteNodeFunc(tp),
			GetNode:    shutdown_get_node.NewGetNodeFunc(tp),
			PutNode:    shutdown_put_node.NewPutNodeFunc(tp),
		},

		// Sql
		Sql: Sql{
			ClearCursor:    sql_clear_cursor.NewClearCursorFunc(tp),
			DeleteAsync:    sql_delete_async.NewDeleteAsyncFunc(tp),
			GetAsync:       sql_get_async.NewGetAsyncFunc(tp),
			GetAsyncStatus: sql_get_async_status.NewGetAsyncStatusFunc(tp),
			Query:          sql_query.NewQueryFunc(tp),
			Translate:      sql_translate.NewTranslateFunc(tp),
		},

		// Ssl
		Ssl: Ssl{
			Certificates: ssl_certificates.NewCertificatesFunc(tp),
		},

		// Synonyms
		Synonyms: Synonyms{
			DeleteSynonym:     synonyms_delete_synonym.NewDeleteSynonymFunc(tp),
			DeleteSynonymRule: synonyms_delete_synonym_rule.NewDeleteSynonymRuleFunc(tp),
			GetSynonym:        synonyms_get_synonym.NewGetSynonymFunc(tp),
			GetSynonymRule:    synonyms_get_synonym_rule.NewGetSynonymRuleFunc(tp),
			GetSynonymsSets:   synonyms_get_synonyms_sets.NewGetSynonymsSetsFunc(tp),
			PutSynonym:        synonyms_put_synonym.NewPutSynonymFunc(tp),
			PutSynonymRule:    synonyms_put_synonym_rule.NewPutSynonymRuleFunc(tp),
		},

		// TextStructure
		TextStructure: TextStructure{
			FindFieldStructure:   text_structure_find_field_structure.NewFindFieldStructureFunc(tp),
			FindMessageStructure: text_structure_find_message_structure.NewFindMessageStructureFunc(tp),
			FindStructure:        text_structure_find_structure.NewFindStructureFunc(tp),
			TestGrokPattern:      text_structure_test_grok_pattern.NewTestGrokPatternFunc(tp),
		},

		// Transform
		Transform: Transform{
			DeleteTransform:      transform_delete_transform.NewDeleteTransformFunc(tp),
			GetNodeStats:         transform_get_node_stats.NewGetNodeStatsFunc(tp),
			GetTransform:         transform_get_transform.NewGetTransformFunc(tp),
			GetTransformStats:    transform_get_transform_stats.NewGetTransformStatsFunc(tp),
			PreviewTransform:     transform_preview_transform.NewPreviewTransformFunc(tp),
			PutTransform:         transform_put_transform.NewPutTransformFunc(tp),
			ResetTransform:       transform_reset_transform.NewResetTransformFunc(tp),
			ScheduleNowTransform: transform_schedule_now_transform.NewScheduleNowTransformFunc(tp),
			StartTransform:       transform_start_transform.NewStartTransformFunc(tp),
			StopTransform:        transform_stop_transform.NewStopTransformFunc(tp),
			UpdateTransform:      transform_update_transform.NewUpdateTransformFunc(tp),
			UpgradeTransforms:    transform_upgrade_transforms.NewUpgradeTransformsFunc(tp),
		},

		// Watcher
		Watcher: Watcher{
			AckWatch:        watcher_ack_watch.NewAckWatchFunc(tp),
			ActivateWatch:   watcher_activate_watch.NewActivateWatchFunc(tp),
			DeactivateWatch: watcher_deactivate_watch.NewDeactivateWatchFunc(tp),
			DeleteWatch:     watcher_delete_watch.NewDeleteWatchFunc(tp),
			ExecuteWatch:    watcher_execute_watch.NewExecuteWatchFunc(tp),
			GetSettings:     watcher_get_settings.NewGetSettingsFunc(tp),
			GetWatch:        watcher_get_watch.NewGetWatchFunc(tp),
			PutWatch:        watcher_put_watch.NewPutWatchFunc(tp),
			QueryWatches:    watcher_query_watches.NewQueryWatchesFunc(tp),
			Start:           watcher_start.NewStartFunc(tp),
			Stats:           watcher_stats.NewStatsFunc(tp),
			Stop:            watcher_stop.NewStopFunc(tp),
			UpdateSettings:  watcher_update_settings.NewUpdateSettingsFunc(tp),
		},

		Bulk:                    core_bulk.NewBulkFunc(tp),
		ClearScroll:             core_clear_scroll.NewClearScrollFunc(tp),
		ClosePointInTime:        core_close_point_in_time.NewClosePointInTimeFunc(tp),
		Count:                   core_count.NewCountFunc(tp),
		Create:                  core_create.NewCreateFunc(tp),
		Delete:                  core_delete.NewDeleteFunc(tp),
		DeleteByQuery:           core_delete_by_query.NewDeleteByQueryFunc(tp),
		DeleteByQueryRethrottle: core_delete_by_query_rethrottle.NewDeleteByQueryRethrottleFunc(tp),
		DeleteScript:            core_delete_script.NewDeleteScriptFunc(tp),
		Exists:                  core_exists.NewExistsFunc(tp),
		ExistsSource:            core_exists_source.NewExistsSourceFunc(tp),
		Explain:                 core_explain.NewExplainFunc(tp),
		FieldCaps:               core_field_caps.NewFieldCapsFunc(tp),
		Get:                     core_get.NewGetFunc(tp),
		GetScript:               core_get_script.NewGetScriptFunc(tp),
		GetScriptContext:        core_get_script_context.NewGetScriptContextFunc(tp),
		GetScriptLanguages:      core_get_script_languages.NewGetScriptLanguagesFunc(tp),
		GetSource:               core_get_source.NewGetSourceFunc(tp),
		HealthReport:            core_health_report.NewHealthReportFunc(tp),
		Index:                   core_index.NewIndexFunc(tp),
		Info:                    core_info.NewInfoFunc(tp),
		KnnSearch:               core_knn_search.NewKnnSearchFunc(tp),
		Mget:                    core_mget.NewMgetFunc(tp),
		Msearch:                 core_msearch.NewMsearchFunc(tp),
		MsearchTemplate:         core_msearch_template.NewMsearchTemplateFunc(tp),
		Mtermvectors:            core_mtermvectors.NewMtermvectorsFunc(tp),
		OpenPointInTime:         core_open_point_in_time.NewOpenPointInTimeFunc(tp),
		Ping:                    core_ping.NewPingFunc(tp),
		PutScript:               core_put_script.NewPutScriptFunc(tp),
		RankEval:                core_rank_eval.NewRankEvalFunc(tp),
		Reindex:                 core_reindex.NewReindexFunc(tp),
		ReindexRethrottle:       core_reindex_rethrottle.NewReindexRethrottleFunc(tp),
		RenderSearchTemplate:    core_render_search_template.NewRenderSearchTemplateFunc(tp),
		ScriptsPainlessExecute:  core_scripts_painless_execute.NewScriptsPainlessExecuteFunc(tp),
		Scroll:                  core_scroll.NewScrollFunc(tp),
		Search:                  core_search.NewSearchFunc(tp),
		SearchMvt:               core_search_mvt.NewSearchMvtFunc(tp),
		SearchShards:            core_search_shards.NewSearchShardsFunc(tp),
		SearchTemplate:          core_search_template.NewSearchTemplateFunc(tp),
		TermsEnum:               core_terms_enum.NewTermsEnumFunc(tp),
		Termvectors:             core_termvectors.NewTermvectorsFunc(tp),
		Update:                  core_update.NewUpdateFunc(tp),
		UpdateByQuery:           core_update_by_query.NewUpdateByQueryFunc(tp),
		UpdateByQueryRethrottle: core_update_by_query_rethrottle.NewUpdateByQueryRethrottleFunc(tp),
	}
}
