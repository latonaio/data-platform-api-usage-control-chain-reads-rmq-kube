package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-usage-control-chain-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-usage-control-chain-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var usageControlChain *[]dpfm_api_output_formatter.UsageControlChain

	for _, fn := range accepter {
		switch fn {
		case "UsageControlChain":
			func() {
				usageControlChain = c.UsageControlChain(mtx, input, output, errs, log)
			}()

		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.UsageControlChainGlobal{
		UsageControlChain: usageControlChain,
	}

	return data
}

func (c *DPFMAPICaller) UsageControlChain(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.UsageControlChain {
	where := fmt.Sprintf("WHERE usageControlChain.UsageControlChain = \"%s\" ", input.UsageControlChain.UsageControlChain)

	if input.UsageControlChain.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND usageControlChain.IsMarkedForDeletion = %v", where, *input.UsageControlChain.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_usage_control_chain_usage_control_chain_data AS usageControlChain
		` + where + ` ORDER BY usageControlChain.IsMarkedForDeletion ASC, usageControlChain.UsageControlChain ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToUsageControlChain(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
