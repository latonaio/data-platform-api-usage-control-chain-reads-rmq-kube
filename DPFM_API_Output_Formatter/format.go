package dpfm_api_output_formatter

import (
	"data-platform-api-usage-control-chain-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToUsageControlChain(rows *sql.Rows) (*[]UsageControlChain, error) {
	defer rows.Close()
	usageControlChain := make([]UsageControlChain, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.UsageControlChain{}

		err := rows.Scan(
			&pm.UsageControlChain,
			&pm.UsageControlLess,
			&pm.Perpetual,
			&pm.Rental,
			&pm.Duration,
			&pm.DurationUnit,
			&pm.ValidityStartDate,
			&pm.ValidityStartTime,
			&pm.ValidityEndDate,
			&pm.ValidityEndTime,
			&pm.DeleteAfterValidityEnd,
			&pm.ServiceLabelRestriction,
			&pm.ApplicationRestriction,
			&pm.PurposeRestriction,
			&pm.BusinessPartnerRoleRestriction,
			&pm.DataStateRestriction,
			&pm.NumberOfUsageRestriction,
			&pm.NumberOfActualUsage,
			&pm.IPAddressRestriction,
			&pm.MACAddressRestriction,
			&pm.ModifyIsAllowed,
			&pm.LocalLoggingIsAllowed,
			&pm.RemoteNotificationIsAllowed,
			&pm.DistributeOnlyIfEncrypted,
			&pm.AttachPolicyWhenDistribute,
			&pm.PostalCode,
			&pm.LocalSubRegion,
			&pm.LocalRegion,
			&pm.Country,
			&pm.GlobalRegion,
			&pm.TimeZone,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &usageControlChain, err
		}

		data := pm
		usageControlChain = append(usageControlChain, UsageControlChain{
			UsageControlChain:              data.UsageControlChain,
			UsageControlLess:               data.UsageControlLess,
			Perpetual:                      data.Perpetual,
			Rental:                         data.Rental,
			Duration:                       data.Duration,
			DurationUnit:                   data.DurationUnit,
			ValidityStartDate:              data.ValidityStartDate,
			ValidityStartTime:              data.ValidityStartTime,
			ValidityEndDate:                data.ValidityEndDate,
			ValidityEndTime:                data.ValidityEndTime,
			DeleteAfterValidityEnd:         data.DeleteAfterValidityEnd,
			ServiceLabelRestriction:        data.ServiceLabelRestriction,
			ApplicationRestriction:         data.ApplicationRestriction,
			PurposeRestriction:             data.PurposeRestriction,
			BusinessPartnerRoleRestriction: data.BusinessPartnerRoleRestriction,
			DataStateRestriction:           data.DataStateRestriction,
			NumberOfUsageRestriction:       data.NumberOfUsageRestriction,
			NumberOfActualUsage:            data.NumberOfActualUsage,
			IPAddressRestriction:           data.IPAddressRestriction,
			MACAddressRestriction:          data.MACAddressRestriction,
			ModifyIsAllowed:                data.ModifyIsAllowed,
			LocalLoggingIsAllowed:          data.LocalLoggingIsAllowed,
			RemoteNotificationIsAllowed:    data.RemoteNotificationIsAllowed,
			DistributeOnlyIfEncrypted:      data.DistributeOnlyIfEncrypted,
			AttachPolicyWhenDistribute:     data.AttachPolicyWhenDistribute,
			PostalCode:                     data.PostalCode,
			LocalSubRegion:                 data.LocalSubRegion,
			LocalRegion:                    data.LocalRegion,
			Country:                        data.Country,
			GlobalRegion:                   data.GlobalRegion,
			TimeZone:                       data.TimeZone,
			CreationDate:                   data.CreationDate,
			CreationTime:                   data.CreationTime,
			LastChangeDate:                 data.LastChangeDate,
			LastChangeTime:                 data.LastChangeTime,
			IsMarkedForDeletion:            data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &usageControlChain, nil
	}

	return &usageControlChain, nil
}
