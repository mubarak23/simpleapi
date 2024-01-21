
  export const processMortgageCardBalance = async (developer: User, ): Promise<MortgageCardBalanceDto> => {
    const connection = await getFreshConnection()
    const projectSubscriptionTransactionRepo = connection.getRepository(ProjectSubscriptionTransaction)
    const MortgageCardRepo = connection.getRepository(MortgageCard)
    const MortgageCardMonthlyBalanceRepo = connection.getRepository(MortgageCardMonthlyBalance)
  
    await ProjectService.isUserADeveloper(developer.id);

    const currentPaymentDate = new Date();
    const currentMonth = currentPaymentDate.getMonth()
    const currentYear = currentPaymentDate.getFullYear()


    const developerMortgageCard = await MortgageCardRepo.findOne({
      where: { userId: developer.id, isUsed: true, isActive: true, isSoftDeleted: false}
    })
    const developerMortgageCardMonthlyBalance = await MortgageCardMonthlyBalanceRepo.findOne({
      where: { developerUserId: developer.id, month: currentMonth, year: currentYear, isSoftDeleted: false}
    })


    if(!developerMortgageCard){
      throw new UnprocessableEntityError('No Active Mortgage Card at the Moment')
    }

    const wallet = await walletService.getCustomerWallet(developerMortgageCard.userId)
    const CurrencyEnum: { [idx: string]: CurrencyToSymbol; } = <any>CurrencyToSymbol;
    const currencySymbol = CurrencyEnum[wallet.currency ?? 'NGN'] || '₦'

    
    if(!developerMortgageCardMonthlyBalance){
      const mortgageCardWithBalance: MortgageCardBalanceDto = {
        pan: developerMortgageCard.pan,
        currency: wallet.currency ?? 'NGN',
        currencySymbol,
        amountMajor: 0.00,
        isUsed: developerMortgageCard.isUsed,
        isActive: developerMortgageCard.isActive,
        isSoftDeleted: developerMortgageCard.isSoftDeleted,
        createdAt: developerMortgageCard.createdAt
      }
      return mortgageCardWithBalance;
    }

  //   const wallet = await walletService.getCustomerWallet(developerMortgageCard.userId)

  

    

    const join = {
      alias: "project_susbscription_transactions",
      leftJoinAndSelect: {
        projectSubscription: "project_susbscription_transactions.projectSubscription",
        project: "project_susbscription_transactions.project",
        developer: "project_susbscription_transactions.developer",
        investor: "project_susbscription_transactions.investor",
      },
    }

  
    const projectSubscriptionTransactions = await projectSubscriptionTransactionRepo.find({
      where: { developerUserId: developerMortgageCard.userId,  isPaid: true, paidStatus: PaymentTransactionStatus.PAID},
      join
    })

    if(!projectSubscriptionTransactions){
      const mortgageCardWithBalance: MortgageCardBalanceDto = {
        pan: developerMortgageCard.pan,
        currency: wallet.currency ?? 'NGN',
        currencySymbol,
        amountMajor: 0.00,
        isUsed: developerMortgageCard.isUsed,
        isActive: developerMortgageCard.isActive,
        isSoftDeleted: developerMortgageCard.isSoftDeleted,
        createdAt: developerMortgageCard.createdAt
      }
      return mortgageCardWithBalance;
    }

    const currentDate = new Date();
    const currentMonth = currentDate.getMonth() + 1;
    const currentYear = currentDate.getFullYear();
    // console.log('currentMonth', currentMonth)
    const filteredTransactions = projectSubscriptionTransactions.filter(
      (transaction) => {
        const nextPaymentDate = new Date(transaction.nextPaymentDate);
        const transactionMonth = nextPaymentDate.getMonth(); 
        const transactionYear = nextPaymentDate.getFullYear();
       //  console.log('transactionMonth', transactionMonth)
        return transactionMonth === currentMonth && transactionYear === currentYear;
      }
    );

    if(filteredTransactions.length === 0){
      const mortgageCardWithBalance: MortgageCardBalanceDto = {
        pan: developerMortgageCard.pan,
        currency: wallet.currency ?? 'NGN',
        currencySymbol,
        amountMajor: 0.00,
        isUsed: developerMortgageCard.isUsed,
        isActive: developerMortgageCard.isActive,
        isSoftDeleted: developerMortgageCard.isSoftDeleted,
        createdAt: developerMortgageCard.createdAt
      }
      return mortgageCardWithBalance;
    }

    const investorUserIds = filteredTransactions.map( investor => investor.investorUserId);

    const UnpaidProjectSubscriptionTransactions = await projectSubscriptionTransactionRepo.find({
      where: { developerUserId: developerMortgageCard.userId, investorUserId: In(investorUserIds),  isPaid: false, paidStatus: PaymentTransactionStatus.UNPAID},
      join
    })

    if(UnpaidProjectSubscriptionTransactions.length === 0) {
      const mortgageCardWithBalance: MortgageCardBalanceDto = {
        pan: developerMortgageCard.pan,
        currency: wallet.currency ?? 'NGN',
        currencySymbol,
        amountMajor: 0.00,
        isUsed: developerMortgageCard.isUsed,
        isActive: developerMortgageCard.isActive,
        isSoftDeleted: developerMortgageCard.isSoftDeleted,
        createdAt: developerMortgageCard.createdAt
      }
      return mortgageCardWithBalance;
    }
  
    const nextPaymentMonth = Utils.nextPaymentDate30days(currentDate.toISOString());

    const filteredUnpaidTransactions = projectSubscriptionTransactions.filter(
      (transaction) => {
        const nextPaymentDate = new Date(transaction.nextPaymentDate);
        const transactionMonth = nextPaymentDate.getMonth(); 
       //  console.log('transactionMonth', transactionMonth)
        return transactionMonth === nextPaymentMonth;
      }
    );

  if(filteredUnpaidTransactions.length === 0){
    const mortgageCardWithBalance: MortgageCardBalanceDto = {
      pan: developerMortgageCard.pan,
      currency: wallet.currency ?? 'NGN',
      currencySymbol,
      amountMajor: 0.00,
      isUsed: developerMortgageCard.isUsed,
      isActive: developerMortgageCard.isActive,
      isSoftDeleted: developerMortgageCard.isSoftDeleted,
      createdAt: developerMortgageCard.createdAt
    }
    return mortgageCardWithBalance;
  }

  const mortgageCardBalanceMinor =
  filteredUnpaidTransactions.reduce(
    (acc, { amountPaidMinor }) => acc + amountPaidMinor,
    0,
  );
  
  const mortgageCardBalanceMajor = (mortgageCardBalanceMinor || 0) / 100;
  
  const mortgageCardWithBalance: MortgageCardBalanceDto = {
    pan: developerMortgageCard.pan,
    currency: wallet.currency ?? 'NGN',
    currencySymbol,
    amountMajor: mortgageCardBalanceMajor,
    isUsed: developerMortgageCard.isUsed,
    isActive: developerMortgageCard.isActive,
    isSoftDeleted: developerMortgageCard.isSoftDeleted,
    createdAt: developerMortgageCard.createdAt
  }
  return mortgageCardWithBalance;

  }
