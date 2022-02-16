// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly.Inputs
{

    public sealed class CheckGroupAlertSettingsTimeBasedEscalationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// After how many minutes after a check starts failing an alert should be send. Possible values are `5`, `10`, `15`, and `30`. Defaults to `5`.
        /// </summary>
        [Input("minutesFailingThreshold")]
        public Input<int>? MinutesFailingThreshold { get; set; }

        public CheckGroupAlertSettingsTimeBasedEscalationGetArgs()
        {
        }
    }
}
